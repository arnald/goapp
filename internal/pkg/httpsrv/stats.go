package httpsrv

import "log"

type sessionStats struct {
	id   string
	sent int
}

func (w *sessionStats) print() {
	log.Printf("session %s has received %d messages\n", w.id, w.sent)
}

func (w *sessionStats) inc() {
	w.sent++
}

func (s *Server) incStats(id string) {
	for i := range s.sessionStats {
		if s.sessionStats[i].id == id {
			s.sessionStats[i].inc()
			return
		}
	}
	s.sessionStats = append(s.sessionStats, newSessionStats(id))
}

func newSessionStats(id string) sessionStats {
	return sessionStats{id: id, sent: 1}
}
