package crt

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Cond replaces wrangler's condition package with a simpler approach
// using standard K8s condition types.
type Cond string

// ConditionStatus holds standard condition info
type ConditionStatus struct {
	Type               string
	Status             metav1.ConditionStatus
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}

// True sets the condition to True status
func (c Cond) True(obj ConditionAccessor) {
	c.Set(obj, string(metav1.ConditionTrue), "", "")
}

// False sets the condition to False status with a reason and message
func (c Cond) False(obj ConditionAccessor, reason, message string) {
	c.Set(obj, string(metav1.ConditionFalse), reason, message)
}

// Unknown sets the condition to Unknown status
func (c Cond) Unknown(obj ConditionAccessor) {
	c.Set(obj, string(metav1.ConditionUnknown), "", "")
}

// IsTrue returns whether the condition is True
func (c Cond) IsTrue(obj ConditionAccessor) bool {
	cond := c.GetCondition(obj)
	return cond != nil && cond.Status == metav1.ConditionTrue
}

// IsFalse returns whether the condition is False
func (c Cond) IsFalse(obj ConditionAccessor) bool {
	cond := c.GetCondition(obj)
	return cond != nil && cond.Status == metav1.ConditionFalse
}

// GetReason returns the reason for the condition
func (c Cond) GetReason(obj ConditionAccessor) string {
	cond := c.GetCondition(obj)
	if cond == nil {
		return ""
	}
	return cond.Reason
}

// GetMessage returns the message for the condition
func (c Cond) GetMessage(obj ConditionAccessor) string {
	cond := c.GetCondition(obj)
	if cond == nil {
		return ""
	}
	return cond.Message
}

// Set sets the condition fields
func (c Cond) Set(obj ConditionAccessor, status, reason, message string) {
	conditions := obj.GetConditions()
	now := metav1.NewTime(time.Now())

	for i, existing := range conditions {
		if existing.Type == string(c) {
			if existing.Status != metav1.ConditionStatus(status) {
				existing.LastTransitionTime = now
			}
			existing.Status = metav1.ConditionStatus(status)
			existing.Reason = reason
			existing.Message = message
			conditions[i] = existing
			obj.SetConditions(conditions)
			return
		}
	}

	conditions = append(conditions, metav1.Condition{
		Type:               string(c),
		Status:             metav1.ConditionStatus(status),
		LastTransitionTime: now,
		Reason:             reason,
		Message:            message,
	})
	obj.SetConditions(conditions)
}

// GetCondition finds the condition by type
func (c Cond) GetCondition(obj ConditionAccessor) *metav1.Condition {
	for _, cond := range obj.GetConditions() {
		if cond.Type == string(c) {
			return &cond
		}
	}
	return nil
}

// ConditionAccessor is an interface for objects that have conditions
type ConditionAccessor interface {
	GetConditions() []metav1.Condition
	SetConditions([]metav1.Condition)
}
