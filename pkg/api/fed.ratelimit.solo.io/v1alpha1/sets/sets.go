// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets



import (
    fed_ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/fed.ratelimit.solo.io/v1alpha1"

    "github.com/rotisserie/eris"
    sksets "github.com/solo-io/skv2/contrib/pkg/sets"
    "github.com/solo-io/skv2/pkg/ezkube"
    "k8s.io/apimachinery/pkg/util/sets"
)

type FederatedRateLimitConfigSet interface {
	// Get the set stored keys
    Keys() sets.String
    // List of resources stored in the set. Pass an optional filter function to filter on the list.
    List(filterResource ... func(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig) bool) []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig
    // Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
    UnsortedList(filterResource ... func(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig) bool) []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig
    // Return the Set as a map of key to resource.
    Map() map[string]*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig
    // Insert a resource into the set.
    Insert(federatedRateLimitConfig ...*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig)
    // Compare the equality of the keys in two sets (not the resources themselves)
    Equal(federatedRateLimitConfigSet FederatedRateLimitConfigSet) bool
    // Check if the set contains a key matching the resource (not the resource itself)
    Has(federatedRateLimitConfig ezkube.ResourceId) bool
    // Delete the key matching the resource
    Delete(federatedRateLimitConfig  ezkube.ResourceId)
    // Return the union with the provided set
    Union(set FederatedRateLimitConfigSet) FederatedRateLimitConfigSet
    // Return the difference with the provided set
    Difference(set FederatedRateLimitConfigSet) FederatedRateLimitConfigSet
    // Return the intersection with the provided set
    Intersection(set FederatedRateLimitConfigSet) FederatedRateLimitConfigSet
    // Find the resource with the given ID
    Find(id ezkube.ResourceId) (*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig, error)
    // Get the length of the set
    Length() int
    // returns the generic implementation of the set
    Generic() sksets.ResourceSet
    // returns the delta between this and and another FederatedRateLimitConfigSet
    Delta(newSet FederatedRateLimitConfigSet) sksets.ResourceDelta
}

func makeGenericFederatedRateLimitConfigSet(federatedRateLimitConfigList []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig) sksets.ResourceSet {
    var genericResources []ezkube.ResourceId
    for _, obj := range federatedRateLimitConfigList {
        genericResources = append(genericResources, obj)
    }
    return sksets.NewResourceSet(genericResources...)
}

type federatedRateLimitConfigSet struct {
    set sksets.ResourceSet
}

func NewFederatedRateLimitConfigSet(federatedRateLimitConfigList ...*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig) FederatedRateLimitConfigSet {
    return &federatedRateLimitConfigSet{set: makeGenericFederatedRateLimitConfigSet(federatedRateLimitConfigList)}
}

func NewFederatedRateLimitConfigSetFromList(federatedRateLimitConfigList *fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfigList) FederatedRateLimitConfigSet {
    list := make([]*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig, 0, len(federatedRateLimitConfigList.Items))
    for idx := range federatedRateLimitConfigList.Items {
        list = append(list, &federatedRateLimitConfigList.Items[idx])
    }
    return &federatedRateLimitConfigSet{set: makeGenericFederatedRateLimitConfigSet(list)}
}

func (s *federatedRateLimitConfigSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
    }
    return s.Generic().Keys()
}

func (s *federatedRateLimitConfigSet) List(filterResource ... func(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig) bool) []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig))
        })
    }

    objs := s.Generic().List(genericFilters...)
    federatedRateLimitConfigList := make([]*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig, 0, len(objs))
    for _, obj := range objs {
        federatedRateLimitConfigList = append(federatedRateLimitConfigList, obj.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig))
    }
    return federatedRateLimitConfigList
}

func (s *federatedRateLimitConfigSet) UnsortedList(filterResource ... func(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig) bool) []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig {
    if s == nil {
        return nil
    }
    var genericFilters []func(ezkube.ResourceId) bool
    for _, filter := range filterResource {
        genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
            return filter(obj.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig))
        })
    }

    var federatedRateLimitConfigList []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig
    for _, obj := range s.Generic().UnsortedList(genericFilters...) {
        federatedRateLimitConfigList = append(federatedRateLimitConfigList, obj.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig))
    }
    return federatedRateLimitConfigList
}

func (s *federatedRateLimitConfigSet) Map() map[string]*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig {
    if s == nil {
        return nil
    }

    newMap := map[string]*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig{}
    for k, v := range s.Generic().Map() {
        newMap[k] = v.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig)
    }
    return newMap
}

func (s *federatedRateLimitConfigSet) Insert(
        federatedRateLimitConfigList ...*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig,
) {
    if s == nil {
        panic("cannot insert into nil set")
    }

    for _, obj := range federatedRateLimitConfigList {
        s.Generic().Insert(obj)
    }
}

func (s *federatedRateLimitConfigSet) Has(federatedRateLimitConfig ezkube.ResourceId) bool {
    if s == nil {
        return false
    }
    return s.Generic().Has(federatedRateLimitConfig)
}

func (s *federatedRateLimitConfigSet) Equal(
        federatedRateLimitConfigSet FederatedRateLimitConfigSet,
) bool {
    if s == nil {
        return federatedRateLimitConfigSet == nil
    }
    return s.Generic().Equal(federatedRateLimitConfigSet.Generic())
}

func (s *federatedRateLimitConfigSet) Delete(FederatedRateLimitConfig ezkube.ResourceId) {
    if s == nil {
        return
    }
    s.Generic().Delete(FederatedRateLimitConfig)
}

func (s *federatedRateLimitConfigSet) Union(set FederatedRateLimitConfigSet) FederatedRateLimitConfigSet {
    if s == nil {
        return set
    }
    return NewFederatedRateLimitConfigSet(append(s.List(), set.List()...)...)
}

func (s *federatedRateLimitConfigSet) Difference(set FederatedRateLimitConfigSet) FederatedRateLimitConfigSet {
    if s == nil {
        return set
    }
    newSet := s.Generic().Difference(set.Generic())
    return &federatedRateLimitConfigSet{set: newSet}
}

func (s *federatedRateLimitConfigSet) Intersection(set FederatedRateLimitConfigSet) FederatedRateLimitConfigSet {
    if s == nil {
        return nil
    }
    newSet := s.Generic().Intersection(set.Generic())
    var federatedRateLimitConfigList []*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig
    for _, obj := range newSet.List() {
        federatedRateLimitConfigList = append(federatedRateLimitConfigList, obj.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig))
    }
    return NewFederatedRateLimitConfigSet(federatedRateLimitConfigList...)
}


func (s *federatedRateLimitConfigSet) Find(id ezkube.ResourceId) (*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig, error) {
    if s == nil {
        return nil, eris.Errorf("empty set, cannot find FederatedRateLimitConfig %v", sksets.Key(id))
    }
	obj, err := s.Generic().Find(&fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig{}, id)
	if err != nil {
		return nil, err
    }

    return obj.(*fed_ratelimit_solo_io_v1alpha1.FederatedRateLimitConfig), nil
}

func (s *federatedRateLimitConfigSet) Length() int {
    if s == nil {
        return 0
    }
    return s.Generic().Length()
}

func (s *federatedRateLimitConfigSet) Generic() sksets.ResourceSet {
    if s == nil {
        return nil
    }
    return s.set
}

func (s *federatedRateLimitConfigSet) Delta(newSet FederatedRateLimitConfigSet) sksets.ResourceDelta {
    if s == nil {
        return sksets.ResourceDelta{
            Inserted: newSet.Generic(),
        }
    }
    return s.Generic().Delta(newSet.Generic())
}
