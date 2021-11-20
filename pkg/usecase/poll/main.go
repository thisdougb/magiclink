package poll

import (
	"log"
)

func (s *Service) GetNextTask() string {

	job, err := s.repo.GetNextTask()
	if err != nil {
		log.Println("poll.GetNextTask():", err.Error())
	}

	return job
}
