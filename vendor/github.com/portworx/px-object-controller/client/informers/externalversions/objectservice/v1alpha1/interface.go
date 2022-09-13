// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/portworx/px-object-controller/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// PXBucketAccesses returns a PXBucketAccessInformer.
	PXBucketAccesses() PXBucketAccessInformer
	// PXBucketClaims returns a PXBucketClaimInformer.
	PXBucketClaims() PXBucketClaimInformer
	// PXBucketClasses returns a PXBucketClassInformer.
	PXBucketClasses() PXBucketClassInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// PXBucketAccesses returns a PXBucketAccessInformer.
func (v *version) PXBucketAccesses() PXBucketAccessInformer {
	return &pXBucketAccessInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PXBucketClaims returns a PXBucketClaimInformer.
func (v *version) PXBucketClaims() PXBucketClaimInformer {
	return &pXBucketClaimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PXBucketClasses returns a PXBucketClassInformer.
func (v *version) PXBucketClasses() PXBucketClassInformer {
	return &pXBucketClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}