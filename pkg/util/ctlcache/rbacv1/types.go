// Package rbacv1 re-exports wrangler-generated rbac/v1 controller cache types
// so that consuming packages no longer need a direct wrangler import.
package rbacv1

import wranglerrbacv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/rbac/v1"

type ClusterRoleCache = wranglerrbacv1.ClusterRoleCache
type ClusterRoleBindingClient = wranglerrbacv1.ClusterRoleBindingClient
type ClusterRoleBindingCache = wranglerrbacv1.ClusterRoleBindingCache
type RoleBindingClient = wranglerrbacv1.RoleBindingClient
type RoleBindingCache = wranglerrbacv1.RoleBindingCache
