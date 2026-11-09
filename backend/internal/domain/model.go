package domain

import (
	"errors"
	"sort"
	"strings"
	"time"
)

type Seller struct {
	ID, Name, Email        string
	NotificationChannelIDs []string
}
type Buyer struct{ ID, Name, Email string }
type OpsTeamMember struct {
	ID, Name, Email        string
	NotificationChannelIDs []string
}
type Product struct {
	ID, SellerID, Name, Category string
	CurrentPrice                 float64
}
type OrderItem struct {
	ID, OrderID, ProductID, ProductName string
	Quantity                            int
	UnitPrice                           float64
}
type Order struct {
	ID, SellerID, BuyerID, Status string
	OrderedAt                     time.Time
	Items                         []OrderItem
}

func (o Order) Total() float64 {
	var n float64
	for _, i := range o.Items {
		n += float64(i.Quantity) * i.UnitPrice
	}
	return n
}

type Reply struct {
	ID, QuestionID, SellerID, Body string
	CreatedAt                      time.Time
}
type Question struct {
	ID, OrderID, BuyerID, OrderItemID, Body, Status, LifecycleStatus string
	BuyerApprovedAt, ResolvedAt                                      *time.Time
	CreatedAt                                                        time.Time
	Reply                                                            *Reply
	PriorityScore                                                    int
	PriorityClass                                                    string
	MatchedUrgencyWords                                              []string
	PriorityDictionaryVersion                                        string
	LastCategorizedAt                                                time.Time
}
type Account struct {
	ID, Email, Password, DisplayName   string
	Roles                              []string
	SellerID, OpsTeamMemberID, BuyerID string
}
type NotificationChannel struct {
	ID, Type, Destination, Label string
	Enabled                      bool
}
type Notifiable struct {
	ID                     string
	ObjectType             string
	ObjectID               string
	NotificationChannelIDs []string
}

type NotifiableFactory struct{}

func NewNotifiableFactory() NotifiableFactory { return NotifiableFactory{} }

func (NotifiableFactory) FromSeller(seller Seller) (Notifiable, error) {
	if seller.ID == "" {
		return Notifiable{}, errors.New("seller id is required")
	}
	return newNotifiable("seller", seller.ID, seller.NotificationChannelIDs), nil
}

func (NotifiableFactory) FromOpsTeamMember(member OpsTeamMember) (Notifiable, error) {
	if member.ID == "" {
		return Notifiable{}, errors.New("operations team member id is required")
	}
	return newNotifiable("opsTeamMember", member.ID, member.NotificationChannelIDs), nil
}

func newNotifiable(objectType, objectID string, channelIDs []string) Notifiable {
	return Notifiable{
		ID:                     "notifiable-" + objectType + "-" + objectID,
		ObjectType:             objectType,
		ObjectID:               objectID,
		NotificationChannelIDs: append([]string(nil), channelIDs...),
	}
}

type NotificationEvent struct {
	ID                                              string
	QuestionIDs                                     []string
	NotifiableID, ChannelID, BatchDate, GroupingKey string
	IncludedPriorityClasses                         []string
	Message                                         string
	CreatedAt                                       time.Time
}

var CriticalWords = []string{"not delivered", "never arrived", "urgent refund", "cancel immediately", "chargeback", "dangerous", "lawsuit", "stolen", "unsafe", "injury", "broken", "missing", "fraud", "scam", "legal"}
var HighWords = []string{"where is my order", "address issue", "payment issue", "wrong item", "replacement", "complaint", "defective", "damaged", "delayed", "expedite", "refund", "urgent", "late"}

func EvaluatePriority(body string, createdAt, now time.Time, orderValue float64) (int, string, []string) {
	lower := strings.ToLower(body)
	matched := []string{}
	keyword := 0
	for _, w := range CriticalWords {
		if strings.Contains(lower, w) {
			matched = append(matched, w)
			keyword = 80
		}
	}
	if keyword == 0 {
		for _, w := range HighWords {
			if strings.Contains(lower, w) {
				matched = append(matched, w)
				keyword = 60
			}
		}
	}
	days := int(now.Sub(createdAt).Hours() / 24)
	if days < 0 {
		days = 0
	}
	waiting := days * 5
	if waiting > 30 {
		waiting = 30
	}
	value := int(orderValue / 100)
	if value > 20 {
		value = 20
	}
	score := keyword + waiting + value
	if score > 100 {
		score = 100
	}
	class := "Low"
	if score >= 80 {
		class = "Critical"
	} else if score >= 60 {
		class = "High Priority"
	} else if score >= 30 {
		class = "Medium"
	}
	sort.Strings(matched)
	return score, class, matched
}
