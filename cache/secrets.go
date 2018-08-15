package cache

import (
	"errors"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// DefaultCacheSyncTimeout represents the default cache sync timeout
const DefaultCacheSyncTimeout = 10 * time.Second

// SecretsCache queries the Kubernetes API and caches secrets
type SecretsCache interface {
	ListSecrets() ([]v1.Secret, error)
}

// New returns a new secret cache store
func New(client *kubernetes.Clientset, namespace string, resyncPeriod time.Duration) SecretsCache {
	synced := make(chan struct{})

	return &secretsCache{
		client:    client,
		hasSynced: synced,
		store:     newSecretCache(client, namespace, resyncPeriod, synced),
	}
}

type secretsCache struct {
	client    *kubernetes.Clientset
	store     cache.Store
	hasSynced <-chan struct{}
}

func newSecretCache(client *kubernetes.Clientset, namespace string, resyncPeriod time.Duration, synced chan struct{}) cache.Store {
	lw := cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return client.CoreV1().Secrets(namespace).List(opts)
		},
		WatchFunc: func(opts metav1.ListOptions) (watch.Interface, error) {
			return client.CoreV1().Secrets(namespace).Watch(opts)
		},
	}

	store, ctr := cache.NewInformer(&lw,
		&v1.Secret{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{})

	go ctr.Run(nil)

	go func(ctr cache.Controller) {
		for {
			if synced == nil {
				break
			}

			if ctr.HasSynced() {
				close(synced)
				synced = nil
				break
			}
		}
	}(ctr)

	return store
}

func (cache secretsCache) ListSecrets() ([]v1.Secret, error) {
	var secrets []v1.Secret

	if err := cache.blockUntilSync(DefaultCacheSyncTimeout); err != nil {
		return secrets, err
	}

	for _, raw := range cache.store.List() {
		secret, ok := raw.(*v1.Secret)
		if !ok {
			continue
		}

		secrets = append(secrets, *secret)
	}

	return secrets, nil
}

// blockUntilSync blocks until the secrets cache is synced
//
// returns nil if all channels are closed
// returns error if the timeout was reached
func (cache secretsCache) blockUntilSync(timeout time.Duration) error {
	if timeout < 0 {
		<-cache.hasSynced
		return nil
	}

	select {
	case <-time.After(timeout):
		return errors.New("timed out waiting for cache to sync")

	case <-cache.hasSynced:
		return nil
	}
}
