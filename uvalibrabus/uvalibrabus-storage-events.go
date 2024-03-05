//
// Events related to storage lifecycle
//

package uvalibrabus

//
// event names
//

var EventObjectCreate = "storage.object.create"     // object created
var EventObjectUpdate = "storage.object.update"     // object updated
var EventObjectDelete = "storage.object.delete"     // object deleted
var EventMetadataUpdate = "storage.metadata.update" // metadata updated
var EventFileCreate = "storage.file.create"         // file created
var EventFileUpdate = "storage.file.update"         // file updated
var EventFileDelete = "storage.file.delete"         // file deleted

//
// corresponding schema for these events
//

type UvaStorageEvent struct {
	VTag string `json:"vtag"` // object vtag
}

//
// end of file
//
