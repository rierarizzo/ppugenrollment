package enrollment

import "time"

type ApplicationModel struct {
	ID        int       `db:"id"`
	Student   int       `db:"student"`
	Project   int       `db:"project"`
	Schedule  int       `db:"schedule"`
	AppliedOn time.Time `db:"applied_on"`
}
