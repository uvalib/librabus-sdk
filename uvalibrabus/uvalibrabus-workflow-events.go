//
// Events related to workflow lifecycle
//

package uvalibrabus

//
// event names
//

var EventWorkCreate = "workflow.work.create"       // work created
var EventWorkDelete = "workflow.work.delete"       // work deleted
var EventWorkPublish = "workflow.work.publish"     // work published
var EventWorkUnpublish = "workflow.work.unpublish" // work unpublished

//
// corresponding schema for these events
//

type UvaWorkflowEvent struct {
}

//
// end of file
//
