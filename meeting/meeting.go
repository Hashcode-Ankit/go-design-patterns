package meeting

import (
	"fmt"
	"time"
)

type Meet struct {
	id    string
	start time.Time
	end   time.Time
	room  int
}

type meetBuilder struct {
	meeting   Meet
	scheduler func(st time.Time, end time.Time) (bool, error)
}

type MeetingStore struct {
	meetings []Meet
}

func newMeetBuilder(id string, start, end time.Time, room int) *meetBuilder {
	return &meetBuilder{
		meeting: Meet{
			id:    id,
			start: start,
			end:   end,
			room:  room,
		},
	}
}

func (m *meetBuilder) SetScheduler(sched func(st time.Time, end time.Time) (bool, error)) {
	m.scheduler = sched
}

func (m *meetBuilder) ScheduleMeeting() (bool, error) {
	return m.scheduler(m.meeting.start, m.meeting.end)

}

func (m *MeetingStore) CheckOverlappingMeeting(start, end time.Time) bool {
	ch := make(chan string)

	for x := range ch {
		fmt.Println(x)
	}
	return true
}

func MeetingStoreBuilder() *MeetingStore {
	return &MeetingStore{
		meetings: []Meet{},
	}
}
