package object_values

import "github.com/google/uuid"

const (
    ATTACHMENT_TYPE_OTHER_NAME   = "Other"
	ATTACHMENT_TYPE_PROFILE_NAME = "Profile"
)

var (
	ATTACHMENT_TYPE_OTHER_ID = uuid.MustParse("d9904896-5a46-47cd-87b7-ed6977b1f079")
	ATTACHMENT_TYPE_PROFILE_ID   = uuid.MustParse("f7444fbe-5c7f-4ff4-88f1-9fbcfd627b03")
)