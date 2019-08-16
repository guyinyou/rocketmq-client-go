package consumer

import (
	"github.com/apache/rocketmq-client-go/primitive"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAllocateByAveragely(t *testing.T) {
	Convey("Given message queues with a starting value", t, func() {
		queues := []*primitive.MessageQueue{
			{
				QueueId: 0,
			},
			{
				QueueId: 1,
			},
			{
				QueueId: 2,
			},
			{
				QueueId: 3,
			},
			{
				QueueId: 4,
			},
			{
				QueueId: 5,
			},
		}

		Convey("When params is empty", func() {
			result := AllocateByAveragely("testGroup", "", queues, []string{"192.168.24.1@default"})
			So(result, ShouldBeNil)

			result = AllocateByAveragely("testGroup", "192.168.24.1@default", nil, []string{"192.168.24.1@default"})
			So(result, ShouldBeNil)

			result = AllocateByAveragely("testGroup", "192.168.24.1@default", queues, nil)
			So(result, ShouldBeNil)
		})

		type testCase struct {
			currentCid    string
			mqAll         []*primitive.MessageQueue
			cidAll        []string
			expectedQueue []*primitive.MessageQueue
		}
		cases := []testCase{
			{
				currentCid: "192.168.24.1@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 0,
					},
					{
						QueueId: 1,
					},
					{
						QueueId: 2,
					},
				},
			},
			{
				currentCid: "192.168.24.2@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 2,
					},
					{
						QueueId: 3,
					},
				},
			},
			{
				currentCid: "192.168.24.2@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 2,
					},
					{
						QueueId: 3,
					},
				},
			},
			{
				currentCid: "192.168.24.4@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 5,
					},
				},
			},
			{
				currentCid:    "192.168.24.7@default",
				mqAll:         queues,
				cidAll:        []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default", "192.168.24.5@default", "192.168.24.6@default", "192.168.24.7@default"},
				expectedQueue: []*primitive.MessageQueue{},
			},
		}

		Convey("the result of AllocateByAveragely should be deep equal expectedQueue", func() {
			for _, value := range cases {
				result := AllocateByAveragely("testGroup", value.currentCid, value.mqAll, value.cidAll)
				So(result, ShouldResemble, value.expectedQueue)
			}
		})
	})
}

func TestAllocateByAveragelyCircle(t *testing.T) {
	Convey("Given message queues with a starting value", t, func() {
		queues := []*primitive.MessageQueue{
			{
				QueueId: 0,
			},
			{
				QueueId: 1,
			},
			{
				QueueId: 2,
			},
			{
				QueueId: 3,
			},
			{
				QueueId: 4,
			},
			{
				QueueId: 5,
			},
		}

		Convey("When params is empty", func() {
			result := AllocateByAveragelyCircle("testGroup", "", queues, []string{"192.168.24.1@default"})
			So(result, ShouldBeNil)

			result = AllocateByAveragelyCircle("testGroup", "192.168.24.1@default", nil, []string{"192.168.24.1@default"})
			So(result, ShouldBeNil)

			result = AllocateByAveragelyCircle("testGroup", "192.168.24.1@default", queues, nil)
			So(result, ShouldBeNil)
		})

		type testCase struct {
			currentCid    string
			mqAll         []*primitive.MessageQueue
			cidAll        []string
			expectedQueue []*primitive.MessageQueue
		}
		cases := []testCase{
			{
				currentCid: "192.168.24.1@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 0,
					},
					{
						QueueId: 2,
					},
					{
						QueueId: 4,
					},
				},
			},
			{
				currentCid: "192.168.24.2@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 1,
					},
					{
						QueueId: 4,
					},
				},
			},
			{
				currentCid: "192.168.24.2@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 1,
					},
					{
						QueueId: 5,
					},
				},
			},
			{
				currentCid: "192.168.24.4@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId: 3,
					},
				},
			},
			{
				currentCid:    "192.168.24.7@default",
				mqAll:         queues,
				cidAll:        []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default", "192.168.24.5@default", "192.168.24.6@default", "192.168.24.7@default"},
				expectedQueue: []*primitive.MessageQueue{},
			},
		}

		Convey("the result of AllocateByAveragelyCircle should be deep equal expectedQueue", func() {
			for _, value := range cases {
				result := AllocateByAveragelyCircle("testGroup", value.currentCid, value.mqAll, value.cidAll)
				So(result, ShouldResemble, value.expectedQueue)
			}
		})
	})
}

func TestAllocateByConfig(t *testing.T) {
	Convey("Given message queues with a starting value", t, func() {
		queues := []*primitive.MessageQueue{
			{
				QueueId: 0,
			},
			{
				QueueId: 1,
			},
			{
				QueueId: 2,
			},
			{
				QueueId: 3,
			},
			{
				QueueId: 4,
			},
			{
				QueueId: 5,
			},
		}

		strategy := AllocateByConfig(queues)
		result := strategy("testGroup", "192.168.24.1@default", queues, []string{"192.168.24.1@default", "192.168.24.2@default"})
		So(result, ShouldResemble, queues)
	})
}

func TestAllocateByMachineRoom(t *testing.T) {
	Convey("Given some consumer IDCs with a starting value", t, func() {
		idcs := []string{"192.168.24.1", "192.168.24.2"}
		strategy := AllocateByMachineRoom(idcs)

		queues := []*primitive.MessageQueue{
			{
				QueueId:    0,
				BrokerName: "192.168.24.1@defaultName",
			},
			{
				QueueId:    1,
				BrokerName: "192.168.24.1@defaultName",
			},
			{
				QueueId:    2,
				BrokerName: "192.168.24.1@defaultName",
			},
			{
				QueueId:    3,
				BrokerName: "192.168.24.2@defaultName",
			},
			{
				QueueId:    4,
				BrokerName: "192.168.24.2@defaultName",
			},
			{
				QueueId:    5,
				BrokerName: "192.168.24.3@defaultName",
			},
		}

		Convey("When params is empty", func() {
			result := strategy("testGroup", "", queues, []string{"192.168.24.1@default"})
			So(result, ShouldBeNil)

			result = strategy("testGroup", "192.168.24.1@default", nil, []string{"192.168.24.1@default"})
			So(result, ShouldBeNil)

			result = strategy("testGroup", "192.168.24.1@default", queues, nil)
			So(result, ShouldBeNil)
		})

		type testCase struct {
			currentCid    string
			mqAll         []*primitive.MessageQueue
			cidAll        []string
			expectedQueue []*primitive.MessageQueue
		}
		cases := []testCase{
			{
				currentCid: "192.168.24.1@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId:    0,
						BrokerName: "192.168.24.1@defaultName",
					},
					{
						QueueId:    1,
						BrokerName: "192.168.24.1@defaultName",
					},
					{
						QueueId:    4,
						BrokerName: "192.168.24.2@defaultName",
					},
				},
			},
			{
				currentCid: "192.168.24.2@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId:    1,
						BrokerName: "192.168.24.1@defaultName",
					},
					{
						QueueId:    4,
						BrokerName: "192.168.24.2@defaultName",
					},
				},
			},
			{
				currentCid: "192.168.24.2@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId:    1,
						BrokerName: "192.168.24.1@defaultName",
					},
				},
			},
			{
				currentCid: "192.168.24.4@default",
				mqAll:      queues,
				cidAll:     []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default"},
				expectedQueue: []*primitive.MessageQueue{
					{
						QueueId:    3,
						BrokerName: "192.168.24.2@defaultName",
					},
				},
			},
			{
				currentCid:    "192.168.24.7@default",
				mqAll:         queues,
				cidAll:        []string{"192.168.24.1@default", "192.168.24.2@default", "192.168.24.3@default", "192.168.24.4@default", "192.168.24.5@default", "192.168.24.6@default", "192.168.24.7@default"},
				expectedQueue: []*primitive.MessageQueue{},
			},
		}

		Convey("the result of AllocateByMachineRoom should be deep equal expectedQueue", func() {
			for _, value := range cases {
				result := strategy("testGroup", value.currentCid, value.mqAll, value.cidAll)
				So(result, ShouldResemble, value.expectedQueue)
			}
		})
	})
}