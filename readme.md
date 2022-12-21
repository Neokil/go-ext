# Go-EXT

This repository contains a collection of utility functions for easier management and more fluent coding on pointers, slices and more.

## Utility functions explained
### Float
```
Compare(f1 float64, f2 float64, precision int) bool
```
`Compare` allows you to check if two floats are equal up to a specified precision.
This is really useful if you have to compare the result of a calculation where floats sometimes
behave strangely.

### Pointers
```
Create[K any](v K) *K
Resolve[K any](v *K) K
ResolveWithFallback[K any](v *K, fallback K) K
```
Using those functions allows you to create or resolve pointers inline without having to do nil-checks
or create local variables. `Create` creates a pointer, while `Resolve` and `ResolveWithFallback` retrieve
the pointer and in case of a `nil` value either return the type default or the fallback value.

### Slices
```
Any[K any](a []K, check func(K) bool) bool
Contains[K comparable](a []K, v K) bool
Convert[K any, V any](a []K, convert func(K) V) []V
ConvertWithErr[K any, V any](a []K, convert func(K) (V, error)) ([]V, error)
Equal[V comparable](a []V, b []V) bool
Intersect[K comparable](a []K, b []K) []K
IntersectCustom[K any, V any, P any](a []K, b []V, compare func(K, V) *P) []P
Reduce[V any](a []V, reduce func(aValue V, bValue V) V) V
Subtract[K comparable](a []K, b []K) []K
Where[K any](a []K, check func(K) bool) []K
```
Those functions abstract the logic of filtering and converting slices to make it easier to grasp.
So instead of having a loop that does an operation the utility-functions allow you to do the same
but in a more readable way.