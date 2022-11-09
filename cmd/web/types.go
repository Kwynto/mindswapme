package main

// Structure for the data template
type tmplHomeData struct {
	isLinks bool
	Link1, Link2, Link3, Link4, Link5, Link6, Link7, Link8, Link9, Link10,
	Link11, Link12, Link13, Link14, Link15, Link16, Link17, Link18, Link19 msmLink
}

type tmplManageData struct {
	User  string
	Hash  string
	Links []msmLink
}
