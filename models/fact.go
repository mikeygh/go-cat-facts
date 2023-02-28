package models

/*

Sample  JSON from  https://cat-fact.herokuapp.com/facts/random

{
    "status": {
        "verified": null,
        "sentCount": 0
    },
    "_id": "62441069f329dce574d40d61",
    "user": "623b08071fbf24e1ca595408",
    "text": "Cats have four paws.",
    "type": "cat",
    "deleted": false,
    "createdAt": "2022-03-30T08:10:17.330Z",
    "updatedAt": "2022-03-30T08:10:17.330Z",
    "__v": 0
}


*/

type Fact struct {
	ID        string `json:"_id"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Type      string `json:"cat"`
	CreatedAt string `json:"createdAt"`
}
