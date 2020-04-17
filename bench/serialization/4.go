package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var dd = `{
		"@context":{
		   "rdfs":"http://www.w3.org/2000/01/rdf-schema#",
		   "schema":"http://schema.org/",
		   "rdfs:subClassOf":{
			  "@type":"@id"
		   },
		   "name":"rdfs:label",
		   "description":"rdfs:comment",
		   "children":{
			  "@reverse":"rdfs:subClassOf"
		   }
		},
		"@type":"rdfs:Class",
		"description":"The most generic type of item.",
		"name":"Thing",
		"@id":"schema:Thing",
		"layer":"core",
		"children":[
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"An action performed by a direct agent and indirect participants upon a direct object...",
			  "name":"Action",
			  "@id":"schema:Action",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of accomplishing something via previous efforts. It is an instantaneous action rather than an ongoing process.",
					"name":"AchieveAction",
					"@id":"schema:AchieveAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AchieveAction",
						  "description":"The act of being defeated in a competitive activity.",
						  "name":"LoseAction",
						  "@id":"schema:LoseAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AchieveAction",
						  "description":"The act of reaching a draw in a competitive activity.",
						  "name":"TieAction",
						  "@id":"schema:TieAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AchieveAction",
						  "description":"The act of achieving victory in a competitive activity.",
						  "name":"WinAction",
						  "@id":"schema:WinAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of forming one's opinion, reaction or sentiment.",
					"name":"AssessAction",
					"@id":"schema:AssessAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AssessAction",
						  "description":"The act of expressing a preference from a set of options or a large or unbounded set of choices/options.",
						  "name":"ChooseAction",
						  "@id":"schema:ChooseAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ChooseAction",
								"description":"The act of expressing a preference from a fixed/finite/structured set of choices/options.",
								"name":"VoteAction",
								"@id":"schema:VoteAction",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AssessAction",
						  "description":"The act of intentionally disregarding the object. An agent ignores an object.",
						  "name":"IgnoreAction",
						  "@id":"schema:IgnoreAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AssessAction",
						  "description":"The act of responding instinctively and emotionally to an object, expressing a sentiment.",
						  "name":"ReactAction",
						  "@id":"schema:ReactAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReactAction",
								"description":"The act of expressing a consistency of opinion with the object...",
								"name":"AgreeAction",
								"@id":"schema:AgreeAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReactAction",
								"description":"The act of expressing a difference of opinion with the object...",
								"name":"DisagreeAction",
								"@id":"schema:DisagreeAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReactAction",
								"description":"The act of expressing a negative sentiment about the object. An agent dislikes an object (a proposition, topic or theme) with participants.",
								"name":"DislikeAction",
								"@id":"schema:DislikeAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReactAction",
								"description":"An agent approves/certifies/likes/supports/sanction an object.",
								"name":"EndorseAction",
								"@id":"schema:EndorseAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReactAction",
								"description":"The act of expressing a positive sentiment about the object. An agent likes an object (a proposition, topic or theme) with participants.",
								"name":"LikeAction",
								"@id":"schema:LikeAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReactAction",
								"description":"The act of expressing a desire about the object. An agent wants an object.",
								"name":"WantAction",
								"@id":"schema:WantAction",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AssessAction",
						  "description":"The act of producing a balanced opinion about the object for an audience...",
						  "name":"ReviewAction",
						  "@id":"schema:ReviewAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of ingesting information/resources/food.",
					"name":"ConsumeAction",
					"@id":"schema:ConsumeAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of swallowing liquids.",
						  "name":"DrinkAction",
						  "@id":"schema:DrinkAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of swallowing solid objects.",
						  "name":"EatAction",
						  "@id":"schema:EatAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of installing an application.",
						  "name":"InstallAction",
						  "@id":"schema:InstallAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of consuming audio content.",
						  "name":"ListenAction",
						  "@id":"schema:ListenAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of consuming written content.",
						  "name":"ReadAction",
						  "@id":"schema:ReadAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of applying an object to its intended purpose.",
						  "name":"UseAction",
						  "@id":"schema:UseAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:UseAction",
								"description":"The act of dressing oneself in clothing.",
								"name":"WearAction",
								"@id":"schema:WearAction",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of consuming static visual content.",
						  "name":"ViewAction",
						  "@id":"schema:ViewAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ConsumeAction",
						  "description":"The act of consuming dynamic/moving visual content.",
						  "name":"WatchAction",
						  "@id":"schema:WatchAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"An agent controls a device or application.",
					"name":"ControlAction",
					"@id":"schema:ControlAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ControlAction",
						  "description":"The act of starting or activating a device or application (e.g...",
						  "name":"ActivateAction",
						  "@id":"schema:ActivateAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ControlAction",
						  "description":"The act of stopping or deactivating a device or application (e...",
						  "name":"DeactivateAction",
						  "@id":"schema:DeactivateAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ControlAction",
						  "description":"The act of resuming a device or application which was formerly paused (e...",
						  "name":"ResumeAction",
						  "@id":"schema:ResumeAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ControlAction",
						  "description":"The act of momentarily pausing a device or application (e.g. pause music playback or pause a timer).",
						  "name":"SuspendAction",
						  "@id":"schema:SuspendAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of deliberately creating/producing/generating/building a result out of the agent.",
					"name":"CreateAction",
					"@id":"schema:CreateAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreateAction",
						  "description":"The act of producing/preparing food.",
						  "name":"CookAction",
						  "@id":"schema:CookAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreateAction",
						  "description":"The act of producing a visual/graphical representation of an object, typically with a pen/pencil and paper as instruments.",
						  "name":"DrawAction",
						  "@id":"schema:DrawAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreateAction",
						  "description":"The act of capturing sound and moving images on film, video, or digitally.",
						  "name":"FilmAction",
						  "@id":"schema:FilmAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreateAction",
						  "description":"The act of producing a painting, typically with paint and canvas as instruments.",
						  "name":"PaintAction",
						  "@id":"schema:PaintAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreateAction",
						  "description":"The act of capturing still images of objects using a camera.",
						  "name":"PhotographAction",
						  "@id":"schema:PhotographAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreateAction",
						  "description":"The act of authoring written creative content.",
						  "name":"WriteAction",
						  "@id":"schema:WriteAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of finding an object.\n\nRelated actions:\n\n\nSearchAction: FindAction is generally lead by a SearchAction, but not necessarily.",
					"name":"FindAction",
					"@id":"schema:FindAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:FindAction",
						  "description":"An agent inspects, determines, investigates, inquires, or examines an object's accuracy, quality, condition, or state.",
						  "name":"CheckAction",
						  "@id":"schema:CheckAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:FindAction",
						  "description":"The act of discovering/finding an object.",
						  "name":"DiscoverAction",
						  "@id":"schema:DiscoverAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:FindAction",
						  "description":"An agent tracks an object for updates.\n\nRelated actions:\n\n\nFollowAction: Unlike FollowAction, TrackAction refers to the interest on the location of innanimates objects...",
						  "name":"TrackAction",
						  "@id":"schema:TrackAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of interacting with another person or organization.",
					"name":"InteractAction",
					"@id":"schema:InteractAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of forming a personal connection with someone (object) mutually/bidirectionally/symmetrically...",
						  "name":"BefriendAction",
						  "@id":"schema:BefriendAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of conveying information to another person via a communication medium (instrument) such as speech, email, or telephone conversation.",
						  "name":"CommunicateAction",
						  "@id":"schema:CommunicateAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of posing a question / favor to someone.\n\nRelated actions:\n\n\nReplyAction: Appears generally as a response to AskAction.",
								"name":"AskAction",
								"@id":"schema:AskAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of an agent communicating (service provider, social media, etc) their arrival by registering/confirming for a previously reserved service (e...",
								"name":"CheckInAction",
								"@id":"schema:CheckInAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of an agent communicating (service provider, social media, etc) their departure of a previously reserved service (e...",
								"name":"CheckOutAction",
								"@id":"schema:CheckOutAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of generating a comment about a subject.",
								"name":"CommentAction",
								"@id":"schema:CommentAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of notifying someone of information pertinent to them, with no expectation of a response.",
								"name":"InformAction",
								"@id":"schema:InformAction",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:InformAction",
									  "description":"The act of notifying someone that a future event/action is going to happen as expected...",
									  "name":"ConfirmAction",
									  "@id":"schema:ConfirmAction",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:InformAction",
									  "description":"The act of notifying an event organizer as to whether you expect to attend the event.",
									  "name":"RsvpAction",
									  "@id":"schema:RsvpAction",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of asking someone to attend an event. Reciprocal of RsvpAction.",
								"name":"InviteAction",
								"@id":"schema:InviteAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of responding to a question/message asked/sent by the object...",
								"name":"ReplyAction",
								"@id":"schema:ReplyAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CommunicateAction",
								"description":"The act of distributing content to people for their amusement or edification.",
								"name":"ShareAction",
								"@id":"schema:ShareAction",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of forming a personal connection with someone/something (object) unidirectionally/asymmetrically to get updates polled from...",
						  "name":"FollowAction",
						  "@id":"schema:FollowAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"An agent joins an event/group with participants/friends at a location...",
						  "name":"JoinAction",
						  "@id":"schema:JoinAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"An agent leaves an event / group with participants/friends at a location...",
						  "name":"LeaveAction",
						  "@id":"schema:LeaveAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of marrying a person.",
						  "name":"MarryAction",
						  "@id":"schema:MarryAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of registering to be a user of a service, product or web page...",
						  "name":"RegisterAction",
						  "@id":"schema:RegisterAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of forming a personal connection with someone/something (object) unidirectionally/asymmetrically to get updates pushed to...",
						  "name":"SubscribeAction",
						  "@id":"schema:SubscribeAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:InteractAction",
						  "description":"The act of un-registering from a service.\n\nRelated actions:\n\n\nRegisterAction: antonym of UnRegisterAction...",
						  "name":"UnRegisterAction",
						  "@id":"schema:UnRegisterAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of an agent relocating to a place.\n\nRelated actions:\n\n\nTransferAction: Unlike TransferAction, the subject of the move is a living Person or Organization rather than an inanimate object.",
					"name":"MoveAction",
					"@id":"schema:MoveAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MoveAction",
						  "description":"The act of arriving at a place. An agent arrives at a destination from a fromLocation, optionally with participants.",
						  "name":"ArriveAction",
						  "@id":"schema:ArriveAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MoveAction",
						  "description":"The act of departing from a place. An agent departs from an fromLocation for a destination, optionally with participants.",
						  "name":"DepartAction",
						  "@id":"schema:DepartAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MoveAction",
						  "description":"The act of traveling from an fromLocation to a destination by a specified mode of transport, optionally with participants.",
						  "name":"TravelAction",
						  "@id":"schema:TravelAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of manipulating/administering/supervising/controlling one or more objects.",
					"name":"OrganizeAction",
					"@id":"schema:OrganizeAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:OrganizeAction",
						  "description":"The act of organizing tasks/objects/events by associating resources to it.",
						  "name":"AllocateAction",
						  "@id":"schema:AllocateAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AllocateAction",
								"description":"The act of committing to/adopting an object.\n\nRelated actions:\n\n\nRejectAction: The antonym of AcceptAction.",
								"name":"AcceptAction",
								"@id":"schema:AcceptAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AllocateAction",
								"description":"The act of allocating an action/event/task to some destination (someone or something).",
								"name":"AssignAction",
								"@id":"schema:AssignAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AllocateAction",
								"description":"The act of granting permission to an object.",
								"name":"AuthorizeAction",
								"@id":"schema:AuthorizeAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AllocateAction",
								"description":"The act of rejecting to/adopting an object.\n\nRelated actions:\n\n\nAcceptAction: The antonym of RejectAction.",
								"name":"RejectAction",
								"@id":"schema:RejectAction",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:OrganizeAction",
						  "description":"The act of registering to an organization/service without the guarantee to receive it...",
						  "name":"ApplyAction",
						  "@id":"schema:ApplyAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:OrganizeAction",
						  "description":"An agent bookmarks/flags/labels/tags/marks an object.",
						  "name":"BookmarkAction",
						  "@id":"schema:BookmarkAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:OrganizeAction",
						  "description":"The act of planning the execution of an event/task/action/reservation/plan to a future date.",
						  "name":"PlanAction",
						  "@id":"schema:PlanAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlanAction",
								"description":"The act of asserting that a future event/action is no longer going to happen...",
								"name":"CancelAction",
								"@id":"schema:CancelAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlanAction",
								"description":"Reserving a concrete object.\n\nRelated actions:\n\n\nScheduleAction: Unlike ScheduleAction, ReserveAction reserves concrete objects (e...",
								"name":"ReserveAction",
								"@id":"schema:ReserveAction",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlanAction",
								"description":"Scheduling future actions, events, or tasks.\n\nRelated actions:\n\n\nReserveAction: Unlike ReserveAction, ScheduleAction allocates future actions (e...",
								"name":"ScheduleAction",
								"@id":"schema:ScheduleAction",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of playing/exercising/training/performing for enjoyment, leisure, recreation, Competition or exercise...",
					"name":"PlayAction",
					"@id":"schema:PlayAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PlayAction",
						  "description":"The act of participating in exertive activity for the purposes of improving health and fitness.",
						  "name":"ExerciseAction",
						  "@id":"schema:ExerciseAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PlayAction",
						  "description":"The act of participating in performance arts.",
						  "name":"PerformAction",
						  "@id":"schema:PerformAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of searching for an object.\n\nRelated actions:\n\n\nFindAction: SearchAction generally leads to a FindAction, but not necessarily.",
					"name":"SearchAction",
					"@id":"schema:SearchAction",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of participating in an exchange of goods and services for monetary compensation...",
					"name":"TradeAction",
					"@id":"schema:TradeAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"The act of giving money to a seller in exchange for goods or services rendered...",
						  "name":"BuyAction",
						  "@id":"schema:BuyAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"The act of providing goods, services, or money without compensation, often for philanthropic reasons.",
						  "name":"DonateAction",
						  "@id":"schema:DonateAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"An agent orders an object/product/service to be delivered/sent.",
						  "name":"OrderAction",
						  "@id":"schema:OrderAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"An agent pays a price to a participant.",
						  "name":"PayAction",
						  "@id":"schema:PayAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"An agent orders a (not yet released) object/product/service to be delivered/sent.",
						  "name":"PreOrderAction",
						  "@id":"schema:PreOrderAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"An agent quotes/estimates/appraises an object/product/service with a price at a location/store.",
						  "name":"QuoteAction",
						  "@id":"schema:QuoteAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"The act of giving money in return for temporary use, but not ownership, of an object such as a vehicle or property...",
						  "name":"RentAction",
						  "@id":"schema:RentAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"The act of taking money from a buyer in exchange for goods or services rendered...",
						  "name":"SellAction",
						  "@id":"schema:SellAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TradeAction",
						  "description":"The act of giving money voluntarily to a beneficiary in recognition of services rendered.",
						  "name":"TipAction",
						  "@id":"schema:TipAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of transferring/moving (abstract or concrete) animate or inanimate objects from one place to another.",
					"name":"TransferAction",
					"@id":"schema:TransferAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of obtaining an object under an agreement to return it at a later date...",
						  "name":"BorrowAction",
						  "@id":"schema:BorrowAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of downloading an object.",
						  "name":"DownloadAction",
						  "@id":"schema:DownloadAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of transferring ownership of an object to a destination...",
						  "name":"GiveAction",
						  "@id":"schema:GiveAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of providing an object under an agreement that it will be returned at a later date...",
						  "name":"LendAction",
						  "@id":"schema:LendAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of transferring money from one place to another place...",
						  "name":"MoneyTransfer",
						  "@id":"schema:MoneyTransfer",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of physically/electronically taking delivery of an object thathas been transferred from an origin to a destination...",
						  "name":"ReceiveAction",
						  "@id":"schema:ReceiveAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of returning to the origin that which was previously received (concrete objects) or taken (ownership).",
						  "name":"ReturnAction",
						  "@id":"schema:ReturnAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of physically/electronically dispatching an object for transfer from an origin to a destination...",
						  "name":"SendAction",
						  "@id":"schema:SendAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:TransferAction",
						  "description":"The act of gaining ownership of an object from an origin. Reciprocal of GiveAction...",
						  "name":"TakeAction",
						  "@id":"schema:TakeAction",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Action",
					"description":"The act of managing by changing/editing the state of the object.",
					"name":"UpdateAction",
					"@id":"schema:UpdateAction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UpdateAction",
						  "description":"The act of editing by adding an object to a collection.",
						  "name":"AddAction",
						  "@id":"schema:AddAction",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AddAction",
								"description":"The act of adding at a specific location in an ordered collection.",
								"name":"InsertAction",
								"@id":"schema:InsertAction",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:InsertAction",
									  "description":"The act of inserting at the end if an ordered collection.",
									  "name":"AppendAction",
									  "@id":"schema:AppendAction",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:InsertAction",
									  "description":"The act of inserting at the beginning if an ordered collection.",
									  "name":"PrependAction",
									  "@id":"schema:PrependAction",
									  "layer":"core"
								   }
								]
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UpdateAction",
						  "description":"The act of editing a recipient by removing one of its objects.",
						  "name":"DeleteAction",
						  "@id":"schema:DeleteAction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UpdateAction",
						  "description":"The act of editing a recipient by replacing an old object with a new object.",
						  "name":"ReplaceAction",
						  "@id":"schema:ReplaceAction",
						  "layer":"core"
					   }
					]
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"The most generic kind of creative work, including books, movies, photographs, software programs, etc.",
			  "name":"CreativeWork",
			  "@id":"schema:CreativeWork",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"An intangible type to be applied to any archive content, carrying with it a set of properties required to describe archival items and collections.",
					"name":"ArchiveComponent",
					"@id":"schema:ArchiveComponent",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"An article, such as a news article or piece of investigative report...",
					"name":"Article",
					"@id":"schema:Article",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"An Article that an external entity has paid to place or to produce to its specifications...",
						  "name":"AdvertiserContentArticle",
						  "@id":"schema:AdvertiserContentArticle",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"A NewsArticle is an article whose content reports news, or provides background context and supporting materials for understanding the news...",
						  "name":"NewsArticle",
						  "@id":"schema:NewsArticle",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:NewsArticle",
								"description":"An AnalysisNewsArticle is a NewsArticle that, while based on factual reporting, incorporates the expertise of the author/producer, offering interpretations and conclusions.",
								"name":"AnalysisNewsArticle",
								"@id":"schema:AnalysisNewsArticle",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:NewsArticle",
								"description":"A NewsArticle expressing an open call by a NewsMediaOrganization asking the public for input, insights, clarifications, anecdotes, documentation, etc...",
								"name":"AskPublicNewsArticle",
								"@id":"schema:AskPublicNewsArticle",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:NewsArticle",
								"description":"A NewsArticle providing historical context, definition and detail on a specific topic (aka \"explainer\" or \"backgrounder\")...",
								"name":"BackgroundNewsArticle",
								"@id":"schema:BackgroundNewsArticle",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:NewsArticle",
								"description":"An OpinionNewsArticle is a NewsArticle that primarily expresses opinions rather than journalistic reporting of news and events...",
								"name":"OpinionNewsArticle",
								"@id":"schema:OpinionNewsArticle",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:NewsArticle",
								"description":"The ReportageNewsArticle type is a subtype of NewsArticle representing\n news articles which are the result of journalistic news reporting conventions...",
								"name":"ReportageNewsArticle",
								"@id":"schema:ReportageNewsArticle",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:NewsArticle",
								"description":"A NewsArticle and CriticReview providing a professional critic's assessment of a service, product, performance, or artistic or literary work.",
								"name":"ReviewNewsArticle",
								"@id":"schema:ReviewNewsArticle",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"A Report generated by governmental or non-governmental organization.",
						  "name":"Report",
						  "@id":"schema:Report",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"An Article whose content is primarily [satirical] in nature, i...",
						  "name":"SatiricalArticle",
						  "@id":"schema:SatiricalArticle",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"A scholarly article.",
						  "name":"ScholarlyArticle",
						  "@id":"schema:ScholarlyArticle",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ScholarlyArticle",
								"description":"A scholarly article in the medical domain.",
								"name":"MedicalScholarlyArticle",
								"@id":"schema:MedicalScholarlyArticle",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"A post to a social media platform, including blog posts, tweets, Facebook posts, etc.",
						  "name":"SocialMediaPosting",
						  "@id":"schema:SocialMediaPosting",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SocialMediaPosting",
								"description":"A blog post.",
								"name":"BlogPosting",
								"@id":"schema:BlogPosting",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:BlogPosting",
									  "description":"A blog post intended to provide a rolling textual coverage of an ongoing event through continuous updates.",
									  "name":"LiveBlogPosting",
									  "@id":"schema:LiveBlogPosting",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SocialMediaPosting",
								"description":"A posting to a discussion forum.",
								"name":"DiscussionForumPosting",
								"@id":"schema:DiscussionForumPosting",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Article",
						  "description":"A technical article - Example: How-to (task) topics, step-by-step, procedural troubleshooting, specifications, etc.",
						  "name":"TechArticle",
						  "@id":"schema:TechArticle",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:TechArticle",
								"description":"Reference documentation for application programming interfaces (APIs).",
								"name":"APIReference",
								"@id":"schema:APIReference",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A collection or bound volume of maps, charts, plates or tables, physical or in media form illustrating any subject.",
					"name":"Atlas",
					"@id":"schema:Atlas",
					"layer":"bib"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A blog.",
					"name":"Blog",
					"@id":"schema:Blog",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A book.",
					"name":"Book",
					"@id":"schema:Book",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Book",
						  "description":"An audiobook.",
						  "name":"Audiobook",
						  "@id":"schema:Audiobook",
						  "layer":"bib"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"One of the sections into which a book is divided. A chapter usually has a section number or a name.",
					"name":"Chapter",
					"@id":"schema:Chapter",
					"layer":"bib"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A Claim in Schema.org represents a specific, factually-oriented claim that could be the itemReviewed in a ClaimReview...",
					"name":"Claim",
					"@id":"schema:Claim",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A short TV or radio program or a segment/part of a program.",
					"name":"Clip",
					"@id":"schema:Clip",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Clip",
						  "description":"A short segment/part of a movie.",
						  "name":"MovieClip",
						  "@id":"schema:MovieClip",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Clip",
						  "description":"A short radio program or a segment/part of a radio program.",
						  "name":"RadioClip",
						  "@id":"schema:RadioClip",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Clip",
						  "description":"A short TV program or a segment/part of a TV program.",
						  "name":"TVClip",
						  "@id":"schema:TVClip",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Clip",
						  "description":"A short segment/part of a video game.",
						  "name":"VideoGameClip",
						  "@id":"schema:VideoGameClip",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Computer programming source code. Example: Full (compile ready) solutions, code snippet samples, scripts, templates.",
					"name":"Code",
					"@id":"schema:Code",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A created collection of Creative Works or other artefacts.",
					"name":"Collection",
					"@id":"schema:Collection",
					"layer":"bib"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"The term \"story\" is any indivisible, re-printable\n unit of a comic, including the interior stories, covers, and backmatter...",
					"name":"ComicStory",
					"@id":"schema:ComicStory",
					"layer":"bib",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ComicStory",
						  "description":"The artwork on the cover of a comic.",
						  "name":"ComicCoverArt",
						  "@id":"schema:ComicCoverArt",
						  "layer":"bib"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A comment on an item - for example, a comment on a blog post. The comment's content is expressed via the text property, and its topic via about, properties shared with all CreativeWorks.",
					"name":"Comment",
					"@id":"schema:Comment",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Comment",
						  "description":"An answer offered to a question; perhaps correct, perhaps opinionated or wrong.",
						  "name":"Answer",
						  "@id":"schema:Answer",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Comment",
						  "description":"A comment that corrects CreativeWork.",
						  "name":"CorrectionComment",
						  "@id":"schema:CorrectionComment",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"One or more messages between organizations or people on a particular topic...",
					"name":"Conversation",
					"@id":"schema:Conversation",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A description of an educational course which may be offered as distinct instances at which take place at different times or take place at different locations, or be offered through different media or modes of study...",
					"name":"Course",
					"@id":"schema:Course",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A media season e.g. tv, radio, video game etc.",
					"name":"CreativeWorkSeason",
					"@id":"schema:CreativeWorkSeason",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeason",
						  "description":"A single season of a podcast. Many podcasts do not break down into separate seasons...",
						  "name":"PodcastSeason",
						  "@id":"schema:PodcastSeason",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeason",
						  "description":"Season dedicated to radio broadcast and associated online delivery.",
						  "name":"RadioSeason",
						  "@id":"schema:RadioSeason",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeason",
						  "description":"Season dedicated to TV broadcast and associated online delivery.",
						  "name":"TVSeason",
						  "@id":"schema:TVSeason",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A CreativeWorkSeries in schema.org is a group of related items, typically but not necessarily of the same kind...",
					"name":"CreativeWorkSeries",
					"@id":"schema:CreativeWorkSeries",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"A series of books. Included books can be indicated with the hasPart property.",
						  "name":"BookSeries",
						  "@id":"schema:BookSeries",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"A series of movies. Included movies can be indicated with the hasPart property.",
						  "name":"MovieSeries",
						  "@id":"schema:MovieSeries",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"A publication in any medium issued in successive parts bearing numerical or chronological designations and intended, such as a magazine, scholarly journal, or newspaper to continue indefinitely...",
						  "name":"Periodical",
						  "@id":"schema:Periodical",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Periodical",
								"description":"A sequential publication of comic stories under a\n unifying title, for example \"The Amazing Spider-Man\" or \"Groo the\n Wanderer\".",
								"name":"ComicSeries",
								"@id":"schema:ComicSeries",
								"layer":"bib"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Periodical",
								"description":"A publication containing information about varied topics that are pertinent to general information, a geographic area, or a specific subject matter (i...",
								"name":"Newspaper",
								"@id":"schema:Newspaper",
								"layer":"bib"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"A podcast is an episodic series of digital audio or video files which a user can download and listen to.",
						  "name":"PodcastSeries",
						  "@id":"schema:PodcastSeries",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"CreativeWorkSeries dedicated to radio broadcast and associated online delivery.",
						  "name":"RadioSeries",
						  "@id":"schema:RadioSeries",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"CreativeWorkSeries dedicated to TV broadcast and associated online delivery.",
						  "name":"TVSeries",
						  "@id":"schema:TVSeries",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CreativeWorkSeries",
						  "description":"A video game series.",
						  "name":"VideoGameSeries",
						  "@id":"schema:VideoGameSeries",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A collection of datasets.",
					"name":"DataCatalog",
					"@id":"schema:DataCatalog",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A body of structured information describing some topic(s) of interest.",
					"name":"Dataset",
					"@id":"schema:Dataset",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Dataset",
						  "description":"A single feed providing structured information about one or more entities or topics.",
						  "name":"DataFeed",
						  "@id":"schema:DataFeed",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DataFeed",
								"description":"A CompleteDataFeed is a DataFeed whose standard representation includes content for every item currently in the feed...",
								"name":"CompleteDataFeed",
								"@id":"schema:CompleteDataFeed",
								"layer":"pending"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A set of defined terms for example a set of categories or a classification scheme, a glossary, dictionary or enumeration.",
					"name":"DefinedTermSet",
					"@id":"schema:DefinedTermSet",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:DefinedTermSet",
						  "description":"A set of Category Code values.",
						  "name":"CategoryCodeSet",
						  "@id":"schema:CategoryCodeSet",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A strategy of regulating the intake of food to achieve or maintain a specific health-related goal.",
					"name":"Diet",
					"@id":"schema:Diet",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"An electronic file or document.",
					"name":"DigitalDocument",
					"@id":"schema:DigitalDocument",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:DigitalDocument",
						  "description":"A file containing a note, primarily for the author.",
						  "name":"NoteDigitalDocument",
						  "@id":"schema:NoteDigitalDocument",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:DigitalDocument",
						  "description":"A file containing slides or used for a presentation.",
						  "name":"PresentationDigitalDocument",
						  "@id":"schema:PresentationDigitalDocument",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:DigitalDocument",
						  "description":"A spreadsheet file.",
						  "name":"SpreadsheetDigitalDocument",
						  "@id":"schema:SpreadsheetDigitalDocument",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:DigitalDocument",
						  "description":"A file composed primarily of text.",
						  "name":"TextDigitalDocument",
						  "@id":"schema:TextDigitalDocument",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A picture or diagram made with a pencil, pen, or crayon rather than paint.",
					"name":"Drawing",
					"@id":"schema:Drawing",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"An educational or occupational credential. A diploma, academic degree, certification, qualification, badge, etc...",
					"name":"EducationalOccupationalCredential",
					"@id":"schema:EducationalOccupationalCredential",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A media episode (e.g. TV, radio, video game) which can be part of a series or season.",
					"name":"Episode",
					"@id":"schema:Episode",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Episode",
						  "description":"A single episode of a podcast series.",
						  "name":"PodcastEpisode",
						  "@id":"schema:PodcastEpisode",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Episode",
						  "description":"A radio episode which can be part of a series or season.",
						  "name":"RadioEpisode",
						  "@id":"schema:RadioEpisode",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Episode",
						  "description":"A TV episode which can be part of a series or season.",
						  "name":"TVEpisode",
						  "@id":"schema:TVEpisode",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Fitness-related activity designed for a specific health-related purpose, including defined exercise routines as well as activity prescribed by a clinician.",
					"name":"ExercisePlan",
					"@id":"schema:ExercisePlan",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"The Game type represents things which are games. These are typically rule-governed recreational activities, e...",
					"name":"Game",
					"@id":"schema:Game",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Game",
						  "description":"A video game is an electronic game that involves human interaction with a user interface to generate visual feedback on a video device.",
						  "name":"VideoGame",
						  "@id":"schema:VideoGame",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Guide is a page or article that recommend specific products or services, or aspects of a thing for a user to consider...",
					"name":"Guide",
					"@id":"schema:Guide",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Instructions that explain how to achieve a result by performing a sequence of steps.",
					"name":"HowTo",
					"@id":"schema:HowTo",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:HowTo",
						  "description":"A recipe. For dietary restrictions covered by the recipe, a few common restrictions are enumerated via suitableForDiet...",
						  "name":"Recipe",
						  "@id":"schema:Recipe",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A direction indicating a single action to do in the instructions for how to achieve a result.",
					"name":"HowToDirection",
					"@id":"schema:HowToDirection",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A sub-grouping of steps in the instructions for how to achieve a result (e...",
					"name":"HowToSection",
					"@id":"schema:HowToSection",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A step in the instructions for how to achieve a result. It is an ordered list with HowToDirection and/or HowToTip items.",
					"name":"HowToStep",
					"@id":"schema:HowToStep",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"An explanation in the instructions for how to achieve a result...",
					"name":"HowToTip",
					"@id":"schema:HowToTip",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A legal document such as an act, decree, bill, etc. (enforceable or not) or a component of a legal act (like an article).",
					"name":"Legislation",
					"@id":"schema:Legislation",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Legislation",
						  "description":"A specific object or file containing a Legislation. Note that the same Legislation can be published in multiple files...",
						  "name":"LegislationObject",
						  "@id":"schema:LegislationObject",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A book, document, or piece of music written by hand rather than typed or printed.",
					"name":"Manuscript",
					"@id":"schema:Manuscript",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A map.",
					"name":"Map",
					"@id":"schema:Map",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A media object, such as an image, video, or audio object embedded in a web page or a downloadable dataset i...",
					"name":"MediaObject",
					"@id":"schema:MediaObject",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MediaObject",
						  "description":"A 3D model represents some kind of 3D content, which may have encodings in one or more MediaObjects...",
						  "name":"3DModel",
						  "@id":"schema:3DModel",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MediaObject",
						  "description":"An audio file.",
						  "name":"AudioObject",
						  "@id":"schema:AudioObject",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MediaObject",
						  "description":"A dataset in downloadable form.",
						  "name":"DataDownload",
						  "@id":"schema:DataDownload",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MediaObject",
						  "description":"An image file.",
						  "name":"ImageObject",
						  "@id":"schema:ImageObject",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ImageObject",
								"description":"An image of a visual machine-readable code such as a barcode or QR code.",
								"name":"Barcode",
								"@id":"schema:Barcode",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MediaObject",
						  "description":"A music video file.",
						  "name":"MusicVideoObject",
						  "@id":"schema:MusicVideoObject",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MediaObject",
						  "description":"A video file.",
						  "name":"VideoObject",
						  "@id":"schema:VideoObject",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A structured representation of food or drink items available from a FoodEstablishment.",
					"name":"Menu",
					"@id":"schema:Menu",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A sub-grouping of food or drink items in a menu. E.g. courses (such as 'Dinner', 'Breakfast', etc...",
					"name":"MenuSection",
					"@id":"schema:MenuSection",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A single message from a sender to one or more organizations or people.",
					"name":"Message",
					"@id":"schema:Message",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Message",
						  "description":"An email message.",
						  "name":"EmailMessage",
						  "@id":"schema:EmailMessage",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A movie.",
					"name":"Movie",
					"@id":"schema:Movie",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A musical composition.",
					"name":"MusicComposition",
					"@id":"schema:MusicComposition",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A collection of music tracks in playlist form.",
					"name":"MusicPlaylist",
					"@id":"schema:MusicPlaylist",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MusicPlaylist",
						  "description":"A collection of music tracks.",
						  "name":"MusicAlbum",
						  "@id":"schema:MusicAlbum",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MusicPlaylist",
						  "description":"A MusicRelease is a specific release of a music album.",
						  "name":"MusicRelease",
						  "@id":"schema:MusicRelease",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A music recording (track), usually a single song.",
					"name":"MusicRecording",
					"@id":"schema:MusicRecording",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A painting.",
					"name":"Painting",
					"@id":"schema:Painting",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A photograph.",
					"name":"Photograph",
					"@id":"schema:Photograph",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A play is a form of literature, usually consisting of dialogue between characters, intended for theatrical performance rather than just reading...",
					"name":"Play",
					"@id":"schema:Play",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A large, usually printed placard, bill, or announcement, often illustrated, that is posted to advertise or publicize something.",
					"name":"Poster",
					"@id":"schema:Poster",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A part of a successively published publication such as a periodical or publication volume, often numbered, usually containing a grouping of works such as articles...",
					"name":"PublicationIssue",
					"@id":"schema:PublicationIssue",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PublicationIssue",
						  "description":"Individual comic issues are serially published as\n part of a larger series...",
						  "name":"ComicIssue",
						  "@id":"schema:ComicIssue",
						  "layer":"bib"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A part of a successively published publication such as a periodical or multi-volume work, often numbered...",
					"name":"PublicationVolume",
					"@id":"schema:PublicationVolume",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A specific question - e.g. from a user seeking answers online, or collected in a Frequently Asked Questions (FAQ) document.",
					"name":"Question",
					"@id":"schema:Question",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A quotation. Often but not necessarily from some written work, attributable to a real world author and - if associated with a fictional character - to any fictional Person...",
					"name":"Quotation",
					"@id":"schema:Quotation",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A review of an item - for example, of a restaurant, movie, or store.",
					"name":"Review",
					"@id":"schema:Review",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Review",
						  "description":"A fact-checking review of claims made (or reported) in some creative work (referenced via itemReviewed).",
						  "name":"ClaimReview",
						  "@id":"schema:ClaimReview",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Review",
						  "description":"A CriticReview is a more specialized form of Review written or published by a source that is recognized for its reviewing activities...",
						  "name":"CriticReview",
						  "@id":"schema:CriticReview",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Review",
						  "description":"An EmployerReview is a review of an Organization regarding its role as an employer, written by a current or former employee of that organization.",
						  "name":"EmployerReview",
						  "@id":"schema:EmployerReview",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Review",
						  "description":"(editorial work in progress, this definition is incomplete and unreviewed)\n A MediaReview is a more specialized form of Review dedicated to the evaluation of media content online, typically in the context of fact-checking and misinformation...",
						  "name":"MediaReview",
						  "@id":"schema:MediaReview",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Review",
						  "description":"Recommendation is a type of Review that suggests or proposes something as the best option or best course of action...",
						  "name":"Recommendation",
						  "@id":"schema:Recommendation",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Review",
						  "description":"A review created by an end-user (e.g. consumer, purchaser, attendee etc...",
						  "name":"UserReview",
						  "@id":"schema:UserReview",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A piece of sculpture.",
					"name":"Sculpture",
					"@id":"schema:Sculpture",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A media season e.g. tv, radio, video game etc.",
					"name":"Season",
					"@id":"schema:Season",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Printed music, as opposed to performed or recorded music.",
					"name":"SheetMusic",
					"@id":"schema:SheetMusic",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Short story or tale. A brief work of literature, usually written in narrative prose.",
					"name":"ShortStory",
					"@id":"schema:ShortStory",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A software application.",
					"name":"SoftwareApplication",
					"@id":"schema:SoftwareApplication",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:SoftwareApplication",
						  "description":"A software application designed specifically to work well on a mobile device such as a telephone.",
						  "name":"MobileApplication",
						  "@id":"schema:MobileApplication",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:SoftwareApplication",
						  "description":"Web applications.",
						  "name":"WebApplication",
						  "@id":"schema:WebApplication",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"Computer programming source code. Example: Full (compile ready) solutions, code snippet samples, scripts, templates.",
					"name":"SoftwareSourceCode",
					"@id":"schema:SoftwareSourceCode",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A SpecialAnnouncement combines a simple date-stamped textual information update\n with contextualized Web links and other structured data...",
					"name":"SpecialAnnouncement",
					"@id":"schema:SpecialAnnouncement",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A thesis or dissertation document submitted in support of candidature for an academic degree or professional qualification.",
					"name":"Thesis",
					"@id":"schema:Thesis",
					"layer":"bib"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A work of art that is primarily visual in character.",
					"name":"VisualArtwork",
					"@id":"schema:VisualArtwork",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:VisualArtwork",
						  "description":"The artwork on the outer surface of a CreativeWork.",
						  "name":"CoverArt",
						  "@id":"schema:CoverArt",
						  "layer":"bib"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"WebContent is a type representing all WebPage, WebSite and WebPageElement content...",
					"name":"WebContent",
					"@id":"schema:WebContent",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebContent",
						  "description":"HealthTopicContent is WebContent that is about some aspect of a health topic, e...",
						  "name":"HealthTopicContent",
						  "@id":"schema:HealthTopicContent",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A web page. Every web page is implicitly assumed to be declared to be of type WebPage, so the various properties about that webpage, such as breadcrumb may be used...",
					"name":"WebPage",
					"@id":"schema:WebPage",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"Web page type: About page.",
						  "name":"AboutPage",
						  "@id":"schema:AboutPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"Web page type: Checkout page.",
						  "name":"CheckoutPage",
						  "@id":"schema:CheckoutPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"Web page type: Collection page.",
						  "name":"CollectionPage",
						  "@id":"schema:CollectionPage",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CollectionPage",
								"description":"Web page type: Media gallery page. A mixed-media page that can contains media such as images, videos, and other multimedia.",
								"name":"MediaGallery",
								"@id":"schema:MediaGallery",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MediaGallery",
									  "description":"Web page type: Image gallery page.",
									  "name":"ImageGallery",
									  "@id":"schema:ImageGallery",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MediaGallery",
									  "description":"Web page type: Video gallery page.",
									  "name":"VideoGallery",
									  "@id":"schema:VideoGallery",
									  "layer":"core"
								   }
								]
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"Web page type: Contact page.",
						  "name":"ContactPage",
						  "@id":"schema:ContactPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"A FAQPage is a WebPage presenting one or more \"Frequently asked questions\" (see also QAPage).",
						  "name":"FAQPage",
						  "@id":"schema:FAQPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"A page devoted to a single item, such as a particular product or hotel.",
						  "name":"ItemPage",
						  "@id":"schema:ItemPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"A web page that provides medical information.",
						  "name":"MedicalWebPage",
						  "@id":"schema:MedicalWebPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"Web page type: Profile page.",
						  "name":"ProfilePage",
						  "@id":"schema:ProfilePage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"A QAPage is a WebPage focussed on a specific Question and its Answer(s), e...",
						  "name":"QAPage",
						  "@id":"schema:QAPage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"A RealEstateListing is a listing that describes one or more real-estate Offers (whose businessFunction is typically to lease out, or to sell)...",
						  "name":"RealEstateListing",
						  "@id":"schema:RealEstateListing",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPage",
						  "description":"Web page type: Search results page.",
						  "name":"SearchResultsPage",
						  "@id":"schema:SearchResultsPage",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A web page element, like a table or an image.",
					"name":"WebPageElement",
					"@id":"schema:WebPageElement",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPageElement",
						  "description":"A navigation element of the page.",
						  "name":"SiteNavigationElement",
						  "@id":"schema:SiteNavigationElement",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPageElement",
						  "description":"A table on a Web page.",
						  "name":"Table",
						  "@id":"schema:Table",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPageElement",
						  "description":"An advertising section of the page.",
						  "name":"WPAdBlock",
						  "@id":"schema:WPAdBlock",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPageElement",
						  "description":"The footer section of the page.",
						  "name":"WPFooter",
						  "@id":"schema:WPFooter",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPageElement",
						  "description":"The header section of the page.",
						  "name":"WPHeader",
						  "@id":"schema:WPHeader",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:WebPageElement",
						  "description":"A sidebar section of the page.",
						  "name":"WPSideBar",
						  "@id":"schema:WPSideBar",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:CreativeWork",
					"description":"A WebSite is a set of related web pages and other items typically served from a single web domain and accessible via URLs.",
					"name":"WebSite",
					"@id":"schema:WebSite",
					"layer":"core"
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"An event happening at a certain time and location, such as a concert, lecture, or festival...",
			  "name":"Event",
			  "@id":"schema:Event",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Business event.",
					"name":"BusinessEvent",
					"@id":"schema:BusinessEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Children's event.",
					"name":"ChildrensEvent",
					"@id":"schema:ChildrensEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Comedy event.",
					"name":"ComedyEvent",
					"@id":"schema:ComedyEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"An instance of a Course which is distinct from other instances because it is offered at a different time or location or through different media or modes of study or to a specific section of students.",
					"name":"CourseInstance",
					"@id":"schema:CourseInstance",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: A social dance.",
					"name":"DanceEvent",
					"@id":"schema:DanceEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"An event involving the delivery of an item.",
					"name":"DeliveryEvent",
					"@id":"schema:DeliveryEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Education event.",
					"name":"EducationEvent",
					"@id":"schema:EducationEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"A series of Events. Included events can relate with the series using the superEvent property...",
					"name":"EventSeries",
					"@id":"schema:EventSeries",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Exhibition event, e.g. at a museum, library, archive, tradeshow, ...",
					"name":"ExhibitionEvent",
					"@id":"schema:ExhibitionEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Festival.",
					"name":"Festival",
					"@id":"schema:Festival",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Food event.",
					"name":"FoodEvent",
					"@id":"schema:FoodEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Literary event.",
					"name":"LiteraryEvent",
					"@id":"schema:LiteraryEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Music event.",
					"name":"MusicEvent",
					"@id":"schema:MusicEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"A PublicationEvent corresponds indifferently to the event of publication for a CreativeWork of any type e...",
					"name":"PublicationEvent",
					"@id":"schema:PublicationEvent",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PublicationEvent",
						  "description":"An over the air or online broadcast event.",
						  "name":"BroadcastEvent",
						  "@id":"schema:BroadcastEvent",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PublicationEvent",
						  "description":"A publication event e.g. catch-up TV or radio podcast, during which a program is available on-demand.",
						  "name":"OnDemandEvent",
						  "@id":"schema:OnDemandEvent",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Sales event.",
					"name":"SaleEvent",
					"@id":"schema:SaleEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"A screening of a movie or other video.",
					"name":"ScreeningEvent",
					"@id":"schema:ScreeningEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Social event.",
					"name":"SocialEvent",
					"@id":"schema:SocialEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Sports event.",
					"name":"SportsEvent",
					"@id":"schema:SportsEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Theater performance.",
					"name":"TheaterEvent",
					"@id":"schema:TheaterEvent",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
					"name":"UserInteraction",
					"@id":"schema:UserInteraction",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserBlocks",
						  "@id":"schema:UserBlocks",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserCheckins",
						  "@id":"schema:UserCheckins",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserComments",
						  "@id":"schema:UserComments",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserDownloads",
						  "@id":"schema:UserDownloads",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserLikes",
						  "@id":"schema:UserLikes",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserPageVisits",
						  "@id":"schema:UserPageVisits",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserPlays",
						  "@id":"schema:UserPlays",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserPlusOnes",
						  "@id":"schema:UserPlusOnes",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:UserInteraction",
						  "description":"UserInteraction and its subtypes is an old way of talking about users interacting with pages...",
						  "name":"UserTweets",
						  "@id":"schema:UserTweets",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Event",
					"description":"Event type: Visual arts event.",
					"name":"VisualArtsEvent",
					"@id":"schema:VisualArtsEvent",
					"layer":"core"
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"A utility class that serves as the umbrella for a number of 'intangible' things such as quantities, structured values, etc.",
			  "name":"Intangible",
			  "@id":"schema:Intangible",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A set of requirements that a must be fulfilled in order to perform an Action.",
					"name":"ActionAccessSpecification",
					"@id":"schema:ActionAccessSpecification",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An intangible item that describes an alignment between a learning resource and a node in an educational framework.",
					"name":"AlignmentObject",
					"@id":"schema:AlignmentObject",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Intended audience for an item, i.e. the group for whom the item was created.",
					"name":"Audience",
					"@id":"schema:Audience",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Audience",
						  "description":"A set of characteristics belonging to businesses, e.g. who compose an item's target audience.",
						  "name":"BusinessAudience",
						  "@id":"schema:BusinessAudience",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Audience",
						  "description":"An EducationalAudience.",
						  "name":"EducationalAudience",
						  "@id":"schema:EducationalAudience",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Audience",
						  "description":"Target audiences for medical web pages. Enumerated type.",
						  "name":"MedicalAudience",
						  "@id":"schema:MedicalAudience",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalAudience",
								"description":"A patient is any person recipient of health care services.",
								"name":"Patient",
								"@id":"schema:Patient",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Audience",
						  "description":"A set of characteristics belonging to people, e.g. who compose an item's target audience.",
						  "name":"PeopleAudience",
						  "@id":"schema:PeopleAudience",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PeopleAudience",
								"description":"A set of characteristics describing parents, who can be interested in viewing some content.",
								"name":"ParentAudience",
								"@id":"schema:ParentAudience",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An entity holding detailed information about the available bed types, e...",
					"name":"BedDetails",
					"@id":"schema:BedDetails",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A brand is a name used by an organization or business person for labeling a product, product group, or similar.",
					"name":"Brand",
					"@id":"schema:Brand",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A unique instance of a BroadcastService on a CableOrSatelliteService lineup.",
					"name":"BroadcastChannel",
					"@id":"schema:BroadcastChannel",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:BroadcastChannel",
						  "description":"A unique instance of a radio BroadcastService on a CableOrSatelliteService lineup.",
						  "name":"RadioChannel",
						  "@id":"schema:RadioChannel",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RadioChannel",
								"description":"A radio channel that uses AM.",
								"name":"AMRadioChannel",
								"@id":"schema:AMRadioChannel",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RadioChannel",
								"description":"A radio channel that uses FM.",
								"name":"FMRadioChannel",
								"@id":"schema:FMRadioChannel",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:BroadcastChannel",
						  "description":"A unique instance of a television BroadcastService on a CableOrSatelliteService lineup.",
						  "name":"TelevisionChannel",
						  "@id":"schema:TelevisionChannel",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"The frequency in MHz and the modulation used for a particular BroadcastService.",
					"name":"BroadcastFrequencySpecification",
					"@id":"schema:BroadcastFrequencySpecification",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A class, also often called a 'Type'; equivalent to rdfs:Class.",
					"name":"Class",
					"@id":"schema:Class",
					"layer":"meta"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"This type covers computer programming languages such as Scheme and Lisp, as well as other language-like computer representations...",
					"name":"ComputerLanguage",
					"@id":"schema:ComputerLanguage",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A single item within a larger data feed.",
					"name":"DataFeedItem",
					"@id":"schema:DataFeedItem",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A word, name, acronym, phrase, etc. with a formal definition. Often used in the context of category or subject classification, glossaries or dictionaries, product or creative work types, etc...",
					"name":"DefinedTerm",
					"@id":"schema:DefinedTerm",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:DefinedTerm",
						  "description":"A Category Code.",
						  "name":"CategoryCode",
						  "@id":"schema:CategoryCode",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:CategoryCode",
								"description":"A code for a medical entity.",
								"name":"MedicalCode",
								"@id":"schema:MedicalCode",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A demand entity represents the public, not necessarily binding, not necessarily exclusive, announcement by an organization or person to seek a certain type of goods or services...",
					"name":"Demand",
					"@id":"schema:Demand",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A permission for a particular person or group to access a particular file.",
					"name":"DigitalDocumentPermission",
					"@id":"schema:DigitalDocumentPermission",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A program offered by an institution which determines the learning progress to achieve an outcome, usually a credential like a degree or certificate...",
					"name":"EducationalOccupationalProgram",
					"@id":"schema:EducationalOccupationalProgram",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOccupationalProgram",
						  "description":"A program with both an educational and employment component. Typically based at a workplace and structured around work-based learning, with the aim of instilling competencies related to an occupation...",
						  "name":"WorkBasedProgram",
						  "@id":"schema:WorkBasedProgram",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An entry point, within some Web-based protocol.",
					"name":"EntryPoint",
					"@id":"schema:EntryPoint",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Lists or enumerations\u2014for example, a list of cuisines or music genres, etc.",
					"name":"Enumeration",
					"@id":"schema:Enumeration",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"The status of an Action.",
						  "name":"ActionStatusType",
						  "@id":"schema:ActionStatusType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ActionStatusType",
								"description":"An in-progress action (e.g, while watching the movie, or driving to a location).",
								"name":"ActiveActionStatus",
								"@id":"schema:ActiveActionStatus",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ActionStatusType",
								"description":"An action that has already taken place.",
								"name":"CompletedActionStatus",
								"@id":"schema:CompletedActionStatus",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ActionStatusType",
								"description":"An action that failed to complete. The action's error property and the HTTP return code contain more information about the failure.",
								"name":"FailedActionStatus",
								"@id":"schema:FailedActionStatus",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ActionStatusType",
								"description":"A description of an action that is supported.",
								"name":"PotentialActionStatus",
								"@id":"schema:PotentialActionStatus",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A type of boarding policy used by an airline.",
						  "name":"BoardingPolicyType",
						  "@id":"schema:BoardingPolicyType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BoardingPolicyType",
								"description":"The airline boards by groups based on check-in time, priority, etc.",
								"name":"GroupBoardingPolicy",
								"@id":"schema:GroupBoardingPolicy",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BoardingPolicyType",
								"description":"The airline boards by zones of the plane.",
								"name":"ZoneBoardingPolicy",
								"@id":"schema:ZoneBoardingPolicy",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"The publication format of the book.",
						  "name":"BookFormatType",
						  "@id":"schema:BookFormatType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BookFormatType",
								"description":"Book format: Audiobook. This is an enumerated value for use with the bookFormat property...",
								"name":"AudiobookFormat",
								"@id":"schema:AudiobookFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BookFormatType",
								"description":"Book format: Ebook.",
								"name":"EBook",
								"@id":"schema:EBook",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BookFormatType",
								"description":"Book format: GraphicNovel. May represent a bound collection of ComicIssue instances.",
								"name":"GraphicNovel",
								"@id":"schema:GraphicNovel",
								"layer":"bib"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BookFormatType",
								"description":"Book format: Hardcover.",
								"name":"Hardcover",
								"@id":"schema:Hardcover",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BookFormatType",
								"description":"Book format: Paperback.",
								"name":"Paperback",
								"@id":"schema:Paperback",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A business entity type is a conceptual entity representing the legal form, the size, the main line of business, the position in the value chain, or any combination thereof, of an organization or business person...",
						  "name":"BusinessEntityType",
						  "@id":"schema:BusinessEntityType",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"The business function specifies the type of activity or access (i...",
						  "name":"BusinessFunction",
						  "@id":"schema:BusinessFunction",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Enumerated options related to a ContactPoint.",
						  "name":"ContactPointOption",
						  "@id":"schema:ContactPointOption",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ContactPointOption",
								"description":"Uses devices to support users with hearing impairments.",
								"name":"HearingImpairedSupported",
								"@id":"schema:HearingImpairedSupported",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ContactPointOption",
								"description":"The associated telephone number is toll free.",
								"name":"TollFree",
								"@id":"schema:TollFree",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"The day of the week, e.g. used to specify to which day the opening hours of an OpeningHoursSpecification refer...",
						  "name":"DayOfWeek",
						  "@id":"schema:DayOfWeek",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Thursday and Saturday.",
								"name":"Friday",
								"@id":"schema:Friday",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Sunday and Tuesday.",
								"name":"Monday",
								"@id":"schema:Monday",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"This stands for any day that is a public holiday; it is a placeholder for all official public holidays in some particular location...",
								"name":"PublicHolidays",
								"@id":"schema:PublicHolidays",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Friday and Sunday.",
								"name":"Saturday",
								"@id":"schema:Saturday",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Saturday and Monday.",
								"name":"Sunday",
								"@id":"schema:Sunday",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Wednesday and Friday.",
								"name":"Thursday",
								"@id":"schema:Thursday",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Monday and Wednesday.",
								"name":"Tuesday",
								"@id":"schema:Tuesday",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DayOfWeek",
								"description":"The day of the week between Tuesday and Thursday.",
								"name":"Wednesday",
								"@id":"schema:Wednesday",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A delivery method is a standardized procedure for transferring the product or service to the destination of fulfillment chosen by the customer...",
						  "name":"DeliveryMethod",
						  "@id":"schema:DeliveryMethod",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DeliveryMethod",
								"description":"A DeliveryMethod in which an item is made available via locker.",
								"name":"LockerDelivery",
								"@id":"schema:LockerDelivery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DeliveryMethod",
								"description":"A DeliveryMethod in which an item is collected on site, e.g. in a store or at a box office.",
								"name":"OnSitePickup",
								"@id":"schema:OnSitePickup",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DeliveryMethod",
								"description":"A private parcel service as the delivery mode available for a certain offer...",
								"name":"ParcelService",
								"@id":"schema:ParcelService",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A type of permission which can be granted for accessing a digital document.",
						  "name":"DigitalDocumentPermissionType",
						  "@id":"schema:DigitalDocumentPermissionType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DigitalDocumentPermissionType",
								"description":"Permission to add comments to the document.",
								"name":"CommentPermission",
								"@id":"schema:CommentPermission",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DigitalDocumentPermissionType",
								"description":"Permission to read or view the document.",
								"name":"ReadPermission",
								"@id":"schema:ReadPermission",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DigitalDocumentPermissionType",
								"description":"Permission to write or edit the document.",
								"name":"WritePermission",
								"@id":"schema:WritePermission",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"An EventAttendanceModeEnumeration value is one of potentially several modes of organising an event, relating to whether it is online or offline.",
						  "name":"EventAttendanceModeEnumeration",
						  "@id":"schema:EventAttendanceModeEnumeration",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventAttendanceModeEnumeration",
								"description":"MixedEventAttendanceMode - an event that is conducted as a combination of both offline and online modes.",
								"name":"MixedEventAttendanceMode",
								"@id":"schema:MixedEventAttendanceMode",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventAttendanceModeEnumeration",
								"description":"OfflineEventAttendanceMode - an event that is primarily conducted offline.",
								"name":"OfflineEventAttendanceMode",
								"@id":"schema:OfflineEventAttendanceMode",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventAttendanceModeEnumeration",
								"description":"OnlineEventAttendanceMode - an event that is primarily conducted online.",
								"name":"OnlineEventAttendanceMode",
								"@id":"schema:OnlineEventAttendanceMode",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"EventStatusType is an enumeration type whose instances represent several states that an Event may be in.",
						  "name":"EventStatusType",
						  "@id":"schema:EventStatusType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventStatusType",
								"description":"The event has been cancelled. If the event has multiple startDate values, all are assumed to be cancelled...",
								"name":"EventCancelled",
								"@id":"schema:EventCancelled",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventStatusType",
								"description":"Indicates that the event was changed to allow online participation...",
								"name":"EventMovedOnline",
								"@id":"schema:EventMovedOnline",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventStatusType",
								"description":"The event has been postponed and no new date has been set. The event's previousStartDate should be set.",
								"name":"EventPostponed",
								"@id":"schema:EventPostponed",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventStatusType",
								"description":"The event has been rescheduled. The event's previousStartDate should be set to the old date and the startDate should be set to the event's new date...",
								"name":"EventRescheduled",
								"@id":"schema:EventRescheduled",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EventStatusType",
								"description":"The event is taking place or has taken place on the startDate as scheduled...",
								"name":"EventScheduled",
								"@id":"schema:EventScheduled",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Indicates whether this game is multi-player, co-op or single-player.",
						  "name":"GamePlayMode",
						  "@id":"schema:GamePlayMode",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GamePlayMode",
								"description":"Play mode: CoOp. Co-operative games, where you play on the same team with friends.",
								"name":"CoOp",
								"@id":"schema:CoOp",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GamePlayMode",
								"description":"Play mode: MultiPlayer. Requiring or allowing multiple human players to play simultaneously.",
								"name":"MultiPlayer",
								"@id":"schema:MultiPlayer",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GamePlayMode",
								"description":"Play mode: SinglePlayer. Which is played by a lone player.",
								"name":"SinglePlayer",
								"@id":"schema:SinglePlayer",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Status of a game server.",
						  "name":"GameServerStatus",
						  "@id":"schema:GameServerStatus",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GameServerStatus",
								"description":"Game server status: OfflinePermanently. Server is offline and not available.",
								"name":"OfflinePermanently",
								"@id":"schema:OfflinePermanently",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GameServerStatus",
								"description":"Game server status: OfflineTemporarily. Server is offline now but it can be online soon.",
								"name":"OfflineTemporarily",
								"@id":"schema:OfflineTemporarily",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GameServerStatus",
								"description":"Game server status: Online. Server is available.",
								"name":"Online",
								"@id":"schema:Online",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GameServerStatus",
								"description":"Game server status: OnlineFull. Server is online but unavailable...",
								"name":"OnlineFull",
								"@id":"schema:OnlineFull",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"An enumeration of genders.",
						  "name":"GenderType",
						  "@id":"schema:GenderType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GenderType",
								"description":"The female gender.",
								"name":"Female",
								"@id":"schema:Female",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GenderType",
								"description":"The male gender.",
								"name":"Male",
								"@id":"schema:Male",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"HealthAspectEnumeration enumerates several aspects of health content online, each of which might be described using hasHealthAspect and HealthTopicContent.",
						  "name":"HealthAspectEnumeration",
						  "@id":"schema:HealthAspectEnumeration",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Content about the benefits and advantages of usage or utilization of topic.",
								"name":"BenefitsHealthAspect",
								"@id":"schema:BenefitsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Information about the causes and main actions that gave rise to the topic.",
								"name":"CausesHealthAspect",
								"@id":"schema:CausesHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Content about contagion mechanisms and contagiousness information over the topic.",
								"name":"ContagiousnessHealthAspect",
								"@id":"schema:ContagiousnessHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Information about how or where to find a topic. Also may contain location data that can be used for where to look for help if the topic is observed.",
								"name":"HowOrWhereHealthAspect",
								"@id":"schema:HowOrWhereHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Information about coping or life related to the topic.",
								"name":"LivingWithHealthAspect",
								"@id":"schema:LivingWithHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Related topics may be treated by a Topic.",
								"name":"MayTreatHealthAspect",
								"@id":"schema:MayTreatHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Content about common misconceptions and myths that are related to a topic.",
								"name":"MisconceptionsHealthAspect",
								"@id":"schema:MisconceptionsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Overview of the content. Contains a summarized view of the topic with the most relevant information for an introduction.",
								"name":"OverviewHealthAspect",
								"@id":"schema:OverviewHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Content about the real life experience of patients or people that have lived a similar experience about the topic...",
								"name":"PatientExperienceHealthAspect",
								"@id":"schema:PatientExperienceHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Information about actions or measures that can be taken to avoid getting the topic or reaching a critical situation related to the topic.",
								"name":"PreventionHealthAspect",
								"@id":"schema:PreventionHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Typical progression and happenings of life course of the topic.",
								"name":"PrognosisHealthAspect",
								"@id":"schema:PrognosisHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Other prominent or relevant topics tied to the main topic.",
								"name":"RelatedTopicsHealthAspect",
								"@id":"schema:RelatedTopicsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Information about the risk factors and possible complications that may follow a topic.",
								"name":"RisksOrComplicationsHealthAspect",
								"@id":"schema:RisksOrComplicationsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Content about how to screen or further filter a topic.",
								"name":"ScreeningHealthAspect",
								"@id":"schema:ScreeningHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Information about questions that may be asked, when to see a professional, measures before seeing a doctor or content about the first consultation.",
								"name":"SeeDoctorHealthAspect",
								"@id":"schema:SeeDoctorHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Self care actions or measures that can be taken to sooth, health or avoid a topic...",
								"name":"SelfCareHealthAspect",
								"@id":"schema:SelfCareHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Side effects that can be observed from the usage of the topic.",
								"name":"SideEffectsHealthAspect",
								"@id":"schema:SideEffectsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Stages that can be observed from a topic.",
								"name":"StagesHealthAspect",
								"@id":"schema:StagesHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Symptoms or related symptoms of a Topic.",
								"name":"SymptomsHealthAspect",
								"@id":"schema:SymptomsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Treatments or related therapies for a Topic.",
								"name":"TreatmentsHealthAspect",
								"@id":"schema:TreatmentsHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Categorization and other types related to a topic.",
								"name":"TypesHealthAspect",
								"@id":"schema:TypesHealthAspect",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAspectEnumeration",
								"description":"Content about how, when, frequency and dosage of a topic.",
								"name":"UsageOrScheduleHealthAspect",
								"@id":"schema:UsageOrScheduleHealthAspect",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A list of possible product availability options.",
						  "name":"ItemAvailability",
						  "@id":"schema:ItemAvailability",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item has been discontinued.",
								"name":"Discontinued",
								"@id":"schema:Discontinued",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item is in stock.",
								"name":"InStock",
								"@id":"schema:InStock",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item is available only at physical locations.",
								"name":"InStoreOnly",
								"@id":"schema:InStoreOnly",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item has limited availability.",
								"name":"LimitedAvailability",
								"@id":"schema:LimitedAvailability",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item is available only online.",
								"name":"OnlineOnly",
								"@id":"schema:OnlineOnly",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item is out of stock.",
								"name":"OutOfStock",
								"@id":"schema:OutOfStock",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item is available for pre-order.",
								"name":"PreOrder",
								"@id":"schema:PreOrder",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item is available for ordering and delivery before general availability.",
								"name":"PreSale",
								"@id":"schema:PreSale",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemAvailability",
								"description":"Indicates that the item has sold out.",
								"name":"SoldOut",
								"@id":"schema:SoldOut",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Enumerated for values for itemListOrder for indicating how an ordered ItemList is organized.",
						  "name":"ItemListOrderType",
						  "@id":"schema:ItemListOrderType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemListOrderType",
								"description":"An ItemList ordered with lower values listed first.",
								"name":"ItemListOrderAscending",
								"@id":"schema:ItemListOrderAscending",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemListOrderType",
								"description":"An ItemList ordered with higher values listed first.",
								"name":"ItemListOrderDescending",
								"@id":"schema:ItemListOrderDescending",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ItemListOrderType",
								"description":"An ItemList ordered with no explicit order.",
								"name":"ItemListUnordered",
								"@id":"schema:ItemListUnordered",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A list of possible statuses for the legal force of a legislation.",
						  "name":"LegalForceStatus",
						  "@id":"schema:LegalForceStatus",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalForceStatus",
								"description":"Indicates that a legislation is in force.",
								"name":"InForce",
								"@id":"schema:InForce",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalForceStatus",
								"description":"Indicates that a legislation is currently not in force.",
								"name":"NotInForce",
								"@id":"schema:NotInForce",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalForceStatus",
								"description":"Indicates that parts of the legislation are in force, and parts are not.",
								"name":"PartiallyInForce",
								"@id":"schema:PartiallyInForce",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A list of possible levels for the legal validity of a legislation.",
						  "name":"LegalValueLevel",
						  "@id":"schema:LegalValueLevel",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalValueLevel",
								"description":"Indicates that the publisher gives some special status to the publication of the document...",
								"name":"AuthoritativeLegalValue",
								"@id":"schema:AuthoritativeLegalValue",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalValueLevel",
								"description":"Indicates a document for which the text is conclusively what the law says and is legally binding...",
								"name":"DefinitiveLegalValue",
								"@id":"schema:DefinitiveLegalValue",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalValueLevel",
								"description":"All the documents published by an official publisher should have at least the legal value level \"OfficialLegalValue\"...",
								"name":"OfficialLegalValue",
								"@id":"schema:OfficialLegalValue",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalValueLevel",
								"description":"Indicates that a document has no particular or special standing (e...",
								"name":"UnofficialLegalValue",
								"@id":"schema:UnofficialLegalValue",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"An enumeration of several kinds of Map.",
						  "name":"MapCategoryType",
						  "@id":"schema:MapCategoryType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MapCategoryType",
								"description":"A parking map.",
								"name":"ParkingMap",
								"@id":"schema:ParkingMap",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MapCategoryType",
								"description":"A seating map.",
								"name":"SeatingMap",
								"@id":"schema:SeatingMap",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MapCategoryType",
								"description":"A transit map.",
								"name":"TransitMap",
								"@id":"schema:TransitMap",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MapCategoryType",
								"description":"A venue map (e.g. for malls, auditoriums, museums, etc.).",
								"name":"VenueMap",
								"@id":"schema:VenueMap",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"(editorial work in progress, this definition is incomplete and unreviewed) MediaManipulationRatingEnumeration classifies a number of ways in which a media item (video, image, audio) can be manipulated, taking into account the context within which they are published or presented.",
						  "name":"MediaManipulationRatingEnumeration",
						  "@id":"schema:MediaManipulationRatingEnumeration",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MediaManipulationRatingEnumeration",
								"description":"AuthenticMediaObject: An unaltered image that is presented in an accurate way.",
								"name":"AuthenticContent",
								"@id":"schema:AuthenticContent",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MediaManipulationRatingEnumeration",
								"description":"MissingContext: ...",
								"name":"MissingContext",
								"@id":"schema:MissingContext",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Enumerations related to health and the practice of medicine: A concept that is used to attribute a quality to another concept, as a qualifier, a collection of items or a listing of all of the elements of a set in medicine practice.",
						  "name":"MedicalEnumeration",
						  "@id":"schema:MedicalEnumeration",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"A class of medical drugs, e.g., statins. Classes can represent general pharmacological class, common mechanisms of action, common physiological effects, etc.",
								"name":"DrugClass",
								"@id":"schema:DrugClass",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"The cost per unit of a medical drug. Note that this type is not meant to represent the price in an offer of a drug for sale; see the Offer type for that...",
								"name":"DrugCost",
								"@id":"schema:DrugCost",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Enumerated categories of medical drug costs.",
								"name":"DrugCostCategory",
								"@id":"schema:DrugCostCategory",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Categories that represent an assessment of the risk of fetal injury due to a drug or pharmaceutical used as directed by the mother during pregnancy.",
								"name":"DrugPregnancyCategory",
								"@id":"schema:DrugPregnancyCategory",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Indicates whether this drug is available by prescription or over-the-counter.",
								"name":"DrugPrescriptionStatus",
								"@id":"schema:DrugPrescriptionStatus",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Classes of agents or pathogens that transmit infectious diseases...",
								"name":"InfectiousAgentClass",
								"@id":"schema:InfectiousAgentClass",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Categories of medical devices, organized by the purpose or intended use of the device.",
								"name":"MedicalDevicePurpose",
								"@id":"schema:MedicalDevicePurpose",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Level of evidence for a medical guideline. Enumerated type.",
								"name":"MedicalEvidenceLevel",
								"@id":"schema:MedicalEvidenceLevel",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Any medical imaging modality typically used for diagnostic purposes...",
								"name":"MedicalImagingTechnique",
								"@id":"schema:MedicalImagingTechnique",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Design models for observational medical studies. Enumerated type.",
								"name":"MedicalObservationalStudyDesign",
								"@id":"schema:MedicalObservationalStudyDesign",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"An enumeration that describes different types of medical procedures.",
								"name":"MedicalProcedureType",
								"@id":"schema:MedicalProcedureType",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Any specific branch of medical science or practice. Medical specialities include clinical specialties that pertain to particular organ systems and their respective disease states, as well as allied health specialties...",
								"name":"MedicalSpecialty",
								"@id":"schema:MedicalSpecialty",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"The status of a medical study. Enumerated type.",
								"name":"MedicalStudyStatus",
								"@id":"schema:MedicalStudyStatus",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Design models for medical trials. Enumerated type.",
								"name":"MedicalTrialDesign",
								"@id":"schema:MedicalTrialDesign",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A trial design in which neither the researcher nor the patient knows the details of the treatment the patient was randomly assigned to.",
									  "name":"DoubleBlindedTrial",
									  "@id":"schema:DoubleBlindedTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"An international trial.",
									  "name":"InternationalTrial",
									  "@id":"schema:InternationalTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A trial that takes place at multiple centers.",
									  "name":"MultiCenterTrial",
									  "@id":"schema:MultiCenterTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A trial design in which the researcher knows the full details of the treatment, and so does the patient.",
									  "name":"OpenTrial",
									  "@id":"schema:OpenTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A placebo-controlled trial design.",
									  "name":"PlaceboControlledTrial",
									  "@id":"schema:PlaceboControlledTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A randomized trial design.",
									  "name":"RandomizedTrial",
									  "@id":"schema:RandomizedTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A trial design in which the researcher knows which treatment the patient was randomly assigned to but the patient does not.",
									  "name":"SingleBlindedTrial",
									  "@id":"schema:SingleBlindedTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A trial that takes place at a single center.",
									  "name":"SingleCenterTrial",
									  "@id":"schema:SingleCenterTrial",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTrialDesign",
									  "description":"A trial design in which neither the researcher, the person administering the therapy nor the patient knows the details of the treatment the patient was randomly assigned to.",
									  "name":"TripleBlindedTrial",
									  "@id":"schema:TripleBlindedTrial",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"Systems of medical practice.",
								"name":"MedicineSystem",
								"@id":"schema:MedicineSystem",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalEnumeration",
								"description":"A type of physical examination of a patient performed by a physician.",
								"name":"PhysicalExam",
								"@id":"schema:PhysicalExam",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"MerchantReturnEnumeration enumerates several kinds of product return policy...",
						  "name":"MerchantReturnEnumeration",
						  "@id":"schema:MerchantReturnEnumeration",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MerchantReturnEnumeration",
								"description":"MerchantReturnFiniteReturnWindow: there is a finite window for product returns.",
								"name":"MerchantReturnFiniteReturnWindow",
								"@id":"schema:MerchantReturnFiniteReturnWindow",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MerchantReturnEnumeration",
								"description":"MerchantReturnNotPermitted: product returns are not permitted.",
								"name":"MerchantReturnNotPermitted",
								"@id":"schema:MerchantReturnNotPermitted",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MerchantReturnEnumeration",
								"description":"MerchantReturnUnlimitedWindow: there is an unlimited window for product returns.",
								"name":"MerchantReturnUnlimitedWindow",
								"@id":"schema:MerchantReturnUnlimitedWindow",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MerchantReturnEnumeration",
								"description":"MerchantReturnUnspecified: a product return policy is not specified here.",
								"name":"MerchantReturnUnspecified",
								"@id":"schema:MerchantReturnUnspecified",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Classification of the album by it's type of content: soundtrack, live album, studio album, etc.",
						  "name":"MusicAlbumProductionType",
						  "@id":"schema:MusicAlbumProductionType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"CompilationAlbum.",
								"name":"CompilationAlbum",
								"@id":"schema:CompilationAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"DJMixAlbum.",
								"name":"DJMixAlbum",
								"@id":"schema:DJMixAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"DemoAlbum.",
								"name":"DemoAlbum",
								"@id":"schema:DemoAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"LiveAlbum.",
								"name":"LiveAlbum",
								"@id":"schema:LiveAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"MixtapeAlbum.",
								"name":"MixtapeAlbum",
								"@id":"schema:MixtapeAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"RemixAlbum.",
								"name":"RemixAlbum",
								"@id":"schema:RemixAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"SoundtrackAlbum.",
								"name":"SoundtrackAlbum",
								"@id":"schema:SoundtrackAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"SpokenWordAlbum.",
								"name":"SpokenWordAlbum",
								"@id":"schema:SpokenWordAlbum",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumProductionType",
								"description":"StudioAlbum.",
								"name":"StudioAlbum",
								"@id":"schema:StudioAlbum",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"The kind of release which this album is: single, EP or album.",
						  "name":"MusicAlbumReleaseType",
						  "@id":"schema:MusicAlbumReleaseType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumReleaseType",
								"description":"AlbumRelease.",
								"name":"AlbumRelease",
								"@id":"schema:AlbumRelease",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumReleaseType",
								"description":"BroadcastRelease.",
								"name":"BroadcastRelease",
								"@id":"schema:BroadcastRelease",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumReleaseType",
								"description":"EPRelease.",
								"name":"EPRelease",
								"@id":"schema:EPRelease",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicAlbumReleaseType",
								"description":"SingleRelease.",
								"name":"SingleRelease",
								"@id":"schema:SingleRelease",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Format of this release (the type of recording media used, ie. compact disc, digital media, LP, etc...",
						  "name":"MusicReleaseFormatType",
						  "@id":"schema:MusicReleaseFormatType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"CDFormat.",
								"name":"CDFormat",
								"@id":"schema:CDFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"CassetteFormat.",
								"name":"CassetteFormat",
								"@id":"schema:CassetteFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"DVDFormat.",
								"name":"DVDFormat",
								"@id":"schema:DVDFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"DigitalAudioTapeFormat.",
								"name":"DigitalAudioTapeFormat",
								"@id":"schema:DigitalAudioTapeFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"DigitalFormat.",
								"name":"DigitalFormat",
								"@id":"schema:DigitalFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"LaserDiscFormat.",
								"name":"LaserDiscFormat",
								"@id":"schema:LaserDiscFormat",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MusicReleaseFormatType",
								"description":"VinylFormat.",
								"name":"VinylFormat",
								"@id":"schema:VinylFormat",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A list of possible conditions for the item.",
						  "name":"OfferItemCondition",
						  "@id":"schema:OfferItemCondition",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OfferItemCondition",
								"description":"Indicates that the item is damaged.",
								"name":"DamagedCondition",
								"@id":"schema:DamagedCondition",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OfferItemCondition",
								"description":"Indicates that the item is new.",
								"name":"NewCondition",
								"@id":"schema:NewCondition",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OfferItemCondition",
								"description":"Indicates that the item is refurbished.",
								"name":"RefurbishedCondition",
								"@id":"schema:RefurbishedCondition",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OfferItemCondition",
								"description":"Indicates that the item is used.",
								"name":"UsedCondition",
								"@id":"schema:UsedCondition",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Enumerated status values for Order.",
						  "name":"OrderStatus",
						  "@id":"schema:OrderStatus",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing cancellation of an order.",
								"name":"OrderCancelled",
								"@id":"schema:OrderCancelled",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing successful delivery of an order.",
								"name":"OrderDelivered",
								"@id":"schema:OrderDelivered",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing that an order is in transit.",
								"name":"OrderInTransit",
								"@id":"schema:OrderInTransit",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing that payment is due on an order.",
								"name":"OrderPaymentDue",
								"@id":"schema:OrderPaymentDue",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing availability of an order for pickup.",
								"name":"OrderPickupAvailable",
								"@id":"schema:OrderPickupAvailable",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing that there is a problem with the order.",
								"name":"OrderProblem",
								"@id":"schema:OrderProblem",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing that an order is being processed.",
								"name":"OrderProcessing",
								"@id":"schema:OrderProcessing",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrderStatus",
								"description":"OrderStatus representing that an order has been returned.",
								"name":"OrderReturned",
								"@id":"schema:OrderReturned",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A payment method is a standardized procedure for transferring the monetary amount for a purchase...",
						  "name":"PaymentMethod",
						  "@id":"schema:PaymentMethod",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PaymentMethod",
								"description":"A payment method using a credit, debit, store or other card to associate the payment with an account.",
								"name":"PaymentCard",
								"@id":"schema:PaymentCard",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:PaymentCard",
									  "description":"A card payment method of a particular brand or name. Used to mark up a particular payment method and/or the financial product/service that supplies the card account...",
									  "name":"CreditCard",
									  "@id":"schema:CreditCard",
									  "layer":"core"
								   }
								]
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A specific payment status. For example, PaymentDue, PaymentComplete, etc.",
						  "name":"PaymentStatusType",
						  "@id":"schema:PaymentStatusType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PaymentStatusType",
								"description":"An automatic payment system is in place and will be used.",
								"name":"PaymentAutomaticallyApplied",
								"@id":"schema:PaymentAutomaticallyApplied",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PaymentStatusType",
								"description":"The payment has been received and processed.",
								"name":"PaymentComplete",
								"@id":"schema:PaymentComplete",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PaymentStatusType",
								"description":"The payee received the payment, but it was declined for some reason.",
								"name":"PaymentDeclined",
								"@id":"schema:PaymentDeclined",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PaymentStatusType",
								"description":"The payment is due, but still within an acceptable time to be received.",
								"name":"PaymentDue",
								"@id":"schema:PaymentDue",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PaymentStatusType",
								"description":"The payment is due and considered late.",
								"name":"PaymentPastDue",
								"@id":"schema:PaymentPastDue",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Categories of physical activity, organized by physiologic classification.",
						  "name":"PhysicalActivityCategory",
						  "@id":"schema:PhysicalActivityCategory",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Physical activity of relatively low intensity that depends primarily on the aerobic energy-generating process; during activity, the aerobic metabolism uses oxygen to adequately meet energy demands during exercise.",
								"name":"AerobicActivity",
								"@id":"schema:AerobicActivity",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Physical activity that is of high-intensity which utilizes the anaerobic metabolism of the body.",
								"name":"AnaerobicActivity",
								"@id":"schema:AnaerobicActivity",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Physical activity that is engaged to help maintain posture and balance.",
								"name":"Balance",
								"@id":"schema:Balance",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Physical activity that is engaged in to improve joint and muscle flexibility.",
								"name":"Flexibility",
								"@id":"schema:Flexibility",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Any physical activity engaged in for recreational purposes. Examples may include ballroom dancing, roller skating, canoeing, fishing, etc.",
								"name":"LeisureTimeActivity",
								"@id":"schema:LeisureTimeActivity",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Any physical activity engaged in for job-related purposes. Examples may include waiting tables, maid service, carrying a mailbag, picking fruits or vegetables, construction work, etc.",
								"name":"OccupationalActivity",
								"@id":"schema:OccupationalActivity",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PhysicalActivityCategory",
								"description":"Physical activity that is engaged in to improve muscle and bone strength...",
								"name":"StrengthTraining",
								"@id":"schema:StrengthTraining",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"ProductReturnEnumeration enumerates several kinds of product return policy...",
						  "name":"ProductReturnEnumeration",
						  "@id":"schema:ProductReturnEnumeration",
						  "layer":"attic",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ProductReturnEnumeration",
								"description":"ProductReturnFiniteReturnWindow: there is a finite window for product returns.",
								"name":"ProductReturnFiniteReturnWindow",
								"@id":"schema:ProductReturnFiniteReturnWindow",
								"layer":"attic"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ProductReturnEnumeration",
								"description":"ProductReturnNotPermitted: product returns are not permitted.",
								"name":"ProductReturnNotPermitted",
								"@id":"schema:ProductReturnNotPermitted",
								"layer":"attic"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ProductReturnEnumeration",
								"description":"ProductReturnUnlimitedWindow: there is an unlimited window for product returns.",
								"name":"ProductReturnUnlimitedWindow",
								"@id":"schema:ProductReturnUnlimitedWindow",
								"layer":"attic"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ProductReturnEnumeration",
								"description":"ProductReturnUnspecified: a product return policy is not specified here.",
								"name":"ProductReturnUnspecified",
								"@id":"schema:ProductReturnUnspecified",
								"layer":"attic"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A predefined value for a product characteristic, e.g. the power cord plug type 'US' or the garment sizes 'S', 'M', 'L', and 'XL'.",
						  "name":"QualitativeValue",
						  "@id":"schema:QualitativeValue",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:QualitativeValue",
								"description":"A type of bed. This is used for indicating the bed or beds available in an accommodation.",
								"name":"BedType",
								"@id":"schema:BedType",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:QualitativeValue",
								"description":"A value indicating a special usage of a car, e.g. commercial rental, driving school, or as a taxi.",
								"name":"CarUsageType",
								"@id":"schema:CarUsageType",
								"layer":"auto"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:QualitativeValue",
								"description":"A value indicating which roadwheels will receive torque.",
								"name":"DriveWheelConfigurationValue",
								"@id":"schema:DriveWheelConfigurationValue",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:QualitativeValue",
								"description":"A value indicating a steering position.",
								"name":"SteeringPositionValue",
								"@id":"schema:SteeringPositionValue",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"RefundTypeEnumeration enumerates several kinds of product return refund types.",
						  "name":"RefundTypeEnumeration",
						  "@id":"schema:RefundTypeEnumeration",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RefundTypeEnumeration",
								"description":"A ExchangeRefund ...",
								"name":"ExchangeRefund",
								"@id":"schema:ExchangeRefund",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RefundTypeEnumeration",
								"description":"A FullRefund ...",
								"name":"FullRefund",
								"@id":"schema:FullRefund",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RefundTypeEnumeration",
								"description":"A StoreCreditRefund ...",
								"name":"StoreCreditRefund",
								"@id":"schema:StoreCreditRefund",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Enumerated status values for Reservation.",
						  "name":"ReservationStatusType",
						  "@id":"schema:ReservationStatusType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReservationStatusType",
								"description":"The status for a previously confirmed reservation that is now cancelled.",
								"name":"ReservationCancelled",
								"@id":"schema:ReservationCancelled",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReservationStatusType",
								"description":"The status of a confirmed reservation.",
								"name":"ReservationConfirmed",
								"@id":"schema:ReservationConfirmed",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReservationStatusType",
								"description":"The status of a reservation on hold pending an update like credit card number or flight changes.",
								"name":"ReservationHold",
								"@id":"schema:ReservationHold",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReservationStatusType",
								"description":"The status of a reservation when a request has been sent, but not confirmed.",
								"name":"ReservationPending",
								"@id":"schema:ReservationPending",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A diet restricted to certain foods or preparations for cultural, religious, health or lifestyle reasons.",
						  "name":"RestrictedDiet",
						  "@id":"schema:RestrictedDiet",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet appropriate for people with diabetes.",
								"name":"DiabeticDiet",
								"@id":"schema:DiabeticDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet exclusive of gluten.",
								"name":"GlutenFreeDiet",
								"@id":"schema:GlutenFreeDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet conforming to Islamic dietary practices.",
								"name":"HalalDiet",
								"@id":"schema:HalalDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet conforming to Hindu dietary practices, in particular, beef-free.",
								"name":"HinduDiet",
								"@id":"schema:HinduDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet conforming to Jewish dietary practices.",
								"name":"KosherDiet",
								"@id":"schema:KosherDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet focused on reduced calorie intake.",
								"name":"LowCalorieDiet",
								"@id":"schema:LowCalorieDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet focused on reduced fat and cholesterol intake.",
								"name":"LowFatDiet",
								"@id":"schema:LowFatDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet appropriate for people with lactose intolerance.",
								"name":"LowLactoseDiet",
								"@id":"schema:LowLactoseDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet focused on reduced sodium intake.",
								"name":"LowSaltDiet",
								"@id":"schema:LowSaltDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet exclusive of all animal products.",
								"name":"VeganDiet",
								"@id":"schema:VeganDiet",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RestrictedDiet",
								"description":"A diet exclusive of animal meat.",
								"name":"VegetarianDiet",
								"@id":"schema:VegetarianDiet",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"ReturnFeesEnumeration expresses policies for return fees.",
						  "name":"ReturnFeesEnumeration",
						  "@id":"schema:ReturnFeesEnumeration",
						  "layer":"pending",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReturnFeesEnumeration",
								"description":"OriginalShippingFees ...",
								"name":"OriginalShippingFees",
								"@id":"schema:OriginalShippingFees",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReturnFeesEnumeration",
								"description":"RestockingFees ...",
								"name":"RestockingFees",
								"@id":"schema:RestockingFees",
								"layer":"pending"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ReturnFeesEnumeration",
								"description":"ReturnShippingFees ...",
								"name":"ReturnShippingFees",
								"@id":"schema:ReturnShippingFees",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"RsvpResponseType is an enumeration type whose instances represent responding to an RSVP request.",
						  "name":"RsvpResponseType",
						  "@id":"schema:RsvpResponseType",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RsvpResponseType",
								"description":"The invitee may or may not attend.",
								"name":"RsvpResponseMaybe",
								"@id":"schema:RsvpResponseMaybe",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RsvpResponseType",
								"description":"The invitee will not attend.",
								"name":"RsvpResponseNo",
								"@id":"schema:RsvpResponseNo",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:RsvpResponseType",
								"description":"The invitee will attend.",
								"name":"RsvpResponseYes",
								"@id":"schema:RsvpResponseYes",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"Any branch of a field in which people typically develop specific expertise, usually after significant study, time, and effort.",
						  "name":"Specialty",
						  "@id":"schema:Specialty",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Enumeration",
						  "description":"A range of of services that will be provided to a customer free of charge in case of a defect or malfunction of a product...",
						  "name":"WarrantyScope",
						  "@id":"schema:WarrantyScope",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A FloorPlan is an explicit representation of a collection of similar accommodations, allowing the provision of common information (room counts, sizes, layout diagrams) and offers for rental or sale...",
					"name":"FloorPlan",
					"@id":"schema:FloorPlan",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Server that provides game interaction in a multiplayer game.",
					"name":"GameServer",
					"@id":"schema:GameServer",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"(Eventually to be defined as) a supertype of GeoShape designed to accommodate definitions from Geo-Spatial best practices.",
					"name":"GeospatialGeometry",
					"@id":"schema:GeospatialGeometry",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A grant, typically financial or otherwise quantifiable, of resources...",
					"name":"Grant",
					"@id":"schema:Grant",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Grant",
						  "description":"A monetary grant.",
						  "name":"MonetaryGrant",
						  "@id":"schema:MonetaryGrant",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A US-style health insurance plan, including PPOs, EPOs, and HMOs.",
					"name":"HealthInsurancePlan",
					"@id":"schema:HealthInsurancePlan",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A description of costs to the patient under a given network or formulary.",
					"name":"HealthPlanCostSharingSpecification",
					"@id":"schema:HealthPlanCostSharingSpecification",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"For a given health insurance plan, the specification for costs and coverage of prescription drugs.",
					"name":"HealthPlanFormulary",
					"@id":"schema:HealthPlanFormulary",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A US-style health insurance plan network.",
					"name":"HealthPlanNetwork",
					"@id":"schema:HealthPlanNetwork",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A statement of the money due for goods or services; a bill.",
					"name":"Invoice",
					"@id":"schema:Invoice",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A list of items of any sort—for example, Top 10 Movies About Weathermen, or Top 100 Party Songs...",
					"name":"ItemList",
					"@id":"schema:ItemList",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ItemList",
						  "description":"A BreadcrumbList is an ItemList consisting of a chain of linked Web pages, typically described using at least their URL and their name, and typically ending with the current page...",
						  "name":"BreadcrumbList",
						  "@id":"schema:BreadcrumbList",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ItemList",
						  "description":"An OfferCatalog is an ItemList that contains related Offers and/or further OfferCatalogs that are offeredBy the same provider.",
						  "name":"OfferCatalog",
						  "@id":"schema:OfferCatalog",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A listing that describes a job opening in a certain organization.",
					"name":"JobPosting",
					"@id":"schema:JobPosting",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Natural languages such as Spanish, Tamil, Hindi, English, etc...",
					"name":"Language",
					"@id":"schema:Language",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An list item, e.g. a step in a checklist or how-to description.",
					"name":"ListItem",
					"@id":"schema:ListItem",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:ListItem",
						  "description":"An item used as either a tool or supply when performing the instructions for how to to achieve a result.",
						  "name":"HowToItem",
						  "@id":"schema:HowToItem",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HowToItem",
								"description":"A supply consumed when performing the instructions for how to achieve a result.",
								"name":"HowToSupply",
								"@id":"schema:HowToSupply",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HowToItem",
								"description":"A tool used (but not consumed) when performing instructions for how to achieve a result.",
								"name":"HowToTool",
								"@id":"schema:HowToTool",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A subscription which allows a user to access media including audio, video, books, etc.",
					"name":"MediaSubscription",
					"@id":"schema:MediaSubscription",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A food or drink item listed in a menu or menu section.",
					"name":"MenuItem",
					"@id":"schema:MenuItem",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A MerchantReturnPolicy provides information about product return policies associated with an Organization or Product.",
					"name":"MerchantReturnPolicy",
					"@id":"schema:MerchantReturnPolicy",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Instances of the class Observation are used to specify observations about an entity (which may or may not be an instance of a StatisticalPopulation), at a particular time...",
					"name":"Observation",
					"@id":"schema:Observation",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A profession, may involve prolonged training and/or a formal qualification.",
					"name":"Occupation",
					"@id":"schema:Occupation",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An offer to transfer some rights to an item or to provide a service \u2014 for example, an offer to sell tickets to an event, to rent the DVD of a movie, to stream a TV show over the internet, to repair a motorcycle, or to loan a book...",
					"name":"Offer",
					"@id":"schema:Offer",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Offer",
						  "description":"When a single product is associated with multiple offers (for example, the same pair of shoes is offered by different merchants), then AggregateOffer can be used...",
						  "name":"AggregateOffer",
						  "@id":"schema:AggregateOffer",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Offer",
						  "description":"An OfferForLease in Schema.org represents an Offer to lease out something, i...",
						  "name":"OfferForLease",
						  "@id":"schema:OfferForLease",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Offer",
						  "description":"An OfferForPurchase in Schema.org represents an Offer to sell something, i...",
						  "name":"OfferForPurchase",
						  "@id":"schema:OfferForPurchase",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"OfferShippingDetails - indicates the kinds of shipping options might be available for an online shopping offer.",
					"name":"OfferShippingDetails",
					"@id":"schema:OfferShippingDetails",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An order is a confirmation of a transaction (a receipt), which can contain multiple line items, each represented by an Offer that has been accepted by the customer.",
					"name":"Order",
					"@id":"schema:Order",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An order item is a line of an order. It includes the quantity and shipping details of a bought offer.",
					"name":"OrderItem",
					"@id":"schema:OrderItem",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"The delivery of a parcel either via the postal service or a commercial service.",
					"name":"ParcelDelivery",
					"@id":"schema:ParcelDelivery",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A permit issued by an organization, e.g. a parking pass.",
					"name":"Permit",
					"@id":"schema:Permit",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Permit",
						  "description":"A permit issued by a government agency.",
						  "name":"GovernmentPermit",
						  "@id":"schema:GovernmentPermit",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A ProductReturnPolicy provides information about product return policies associated with an Organization or Product.",
					"name":"ProductReturnPolicy",
					"@id":"schema:ProductReturnPolicy",
					"layer":"attic"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Used to describe membership in a loyalty programs (e.g. \"StarAliance\"), traveler clubs (e...",
					"name":"ProgramMembership",
					"@id":"schema:ProgramMembership",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A property, used to indicate attributes and relationships of some Thing; equivalent to rdf:Property.",
					"name":"Property",
					"@id":"schema:Property",
					"layer":"meta"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A Property value specification.",
					"name":"PropertyValueSpecification",
					"@id":"schema:PropertyValueSpecification",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Quantities such as distance, time, mass, weight, etc. Particular instances of say Mass are entities like '3 Kg' or '4 milligrams'.",
					"name":"Quantity",
					"@id":"schema:Quantity",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Quantity",
						  "description":"Properties that take Distances as values are of the form '<Number> <Length unit of measure>'...",
						  "name":"Distance",
						  "@id":"schema:Distance",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Quantity",
						  "description":"Quantity: Duration (use ISO 8601 duration format).",
						  "name":"Duration",
						  "@id":"schema:Duration",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Quantity",
						  "description":"Properties that take Energy as values are of the form '<Number> <Energy unit of measure>'.",
						  "name":"Energy",
						  "@id":"schema:Energy",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Quantity",
						  "description":"Properties that take Mass as values are of the form '<Number> <Mass unit of measure>'...",
						  "name":"Mass",
						  "@id":"schema:Mass",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A rating is an evaluation on a numeric scale, such as 1 to 5 stars.",
					"name":"Rating",
					"@id":"schema:Rating",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Rating",
						  "description":"The average rating based on multiple ratings or reviews.",
						  "name":"AggregateRating",
						  "@id":"schema:AggregateRating",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AggregateRating",
								"description":"An aggregate rating of an Organization related to its role as an employer.",
								"name":"EmployerAggregateRating",
								"@id":"schema:EmployerAggregateRating",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Rating",
						  "description":"An EndorsementRating is a rating that expresses some level of endorsement, for example inclusion in a \"critic's pick\" blog, a\n\"Like\" or \"+1\" on a social network...",
						  "name":"EndorsementRating",
						  "@id":"schema:EndorsementRating",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Describes a reservation for travel, dining or an event. Some reservations require tickets...",
					"name":"Reservation",
					"@id":"schema:Reservation",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for bus travel. \n\nNote: This type is for information about actual reservations, e...",
						  "name":"BusReservation",
						  "@id":"schema:BusReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for an event like a concert, sporting event, or lecture...",
						  "name":"EventReservation",
						  "@id":"schema:EventReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for air travel.\n\nNote: This type is for information about actual reservations, e...",
						  "name":"FlightReservation",
						  "@id":"schema:FlightReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation to dine at a food-related business.\n\nNote: This type is for information about actual reservations, e...",
						  "name":"FoodEstablishmentReservation",
						  "@id":"schema:FoodEstablishmentReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for lodging at a hotel, motel, inn, etc.\n\nNote: This type is for information about actual reservations, e...",
						  "name":"LodgingReservation",
						  "@id":"schema:LodgingReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for a rental car.\n\nNote: This type is for information about actual reservations, e...",
						  "name":"RentalCarReservation",
						  "@id":"schema:RentalCarReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A group of multiple reservations with common values for all sub-reservations.",
						  "name":"ReservationPackage",
						  "@id":"schema:ReservationPackage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for a taxi.\n\nNote: This type is for information about actual reservations, e...",
						  "name":"TaxiReservation",
						  "@id":"schema:TaxiReservation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Reservation",
						  "description":"A reservation for train travel.\n\nNote: This type is for information about actual reservations, e...",
						  "name":"TrainReservation",
						  "@id":"schema:TrainReservation",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Represents additional information about a relationship or property...",
					"name":"Role",
					"@id":"schema:Role",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Role",
						  "description":"A Role that represents a Web link e.g. as expressed via the 'url' property...",
						  "name":"LinkRole",
						  "@id":"schema:LinkRole",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Role",
						  "description":"A subclass of Role used to describe roles within organizations.",
						  "name":"OrganizationRole",
						  "@id":"schema:OrganizationRole",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:OrganizationRole",
								"description":"A subclass of OrganizationRole used to describe employee relationships.",
								"name":"EmployeeRole",
								"@id":"schema:EmployeeRole",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Role",
						  "description":"A PerformanceRole is a Role that some entity places with regard to a theatrical performance, e...",
						  "name":"PerformanceRole",
						  "@id":"schema:PerformanceRole",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A schedule defines a repeating time period used to describe a regularly occurring Event...",
					"name":"Schedule",
					"@id":"schema:Schedule",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Used to describe a seat, such as a reserved seat in an event reservation.",
					"name":"Seat",
					"@id":"schema:Seat",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A Series in schema.org is a group of related items, typically but not necessarily of the same kind...",
					"name":"Series",
					"@id":"schema:Series",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A service provided by an organization, e.g. delivery service, print services, etc.",
					"name":"Service",
					"@id":"schema:Service",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A delivery service through which content is provided via broadcast over the air or online.",
						  "name":"BroadcastService",
						  "@id":"schema:BroadcastService",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BroadcastService",
								"description":"A delivery service through which radio content is provided via broadcast over the air or online.",
								"name":"RadioBroadcastService",
								"@id":"schema:RadioBroadcastService",
								"layer":"pending"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A service which provides access to media programming like TV or radio...",
						  "name":"CableOrSatelliteService",
						  "@id":"schema:CableOrSatelliteService",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A product provided to consumers and businesses by financial institutions such as banks, insurance companies, brokerage firms, consumer finance companies, and investment companies which comprise the financial services industry.",
						  "name":"FinancialProduct",
						  "@id":"schema:FinancialProduct",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialProduct",
								"description":"A product or service offered by a bank whereby one may deposit, withdraw or transfer money and in some cases be paid interest.",
								"name":"BankAccount",
								"@id":"schema:BankAccount",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:BankAccount",
									  "description":"A type of Bank Account with a main purpose of depositing funds to gain interest or other benefits.",
									  "name":"DepositAccount",
									  "@id":"schema:DepositAccount",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialProduct",
								"description":"A service to convert funds from one currency to another currency.",
								"name":"CurrencyConversionService",
								"@id":"schema:CurrencyConversionService",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialProduct",
								"description":"A type of financial product that typically requires the client to transfer funds to a financial service in return for potential beneficial financial return.",
								"name":"InvestmentOrDeposit",
								"@id":"schema:InvestmentOrDeposit",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:InvestmentOrDeposit",
									  "description":"An account that allows an investor to deposit funds and place investment orders with a licensed broker or brokerage firm.",
									  "name":"BrokerageAccount",
									  "@id":"schema:BrokerageAccount",
									  "layer":"pending"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:InvestmentOrDeposit",
									  "description":"A company or fund that gathers capital from a number of investors to create a pool of money that is then re-invested into stocks, bonds and other assets.",
									  "name":"InvestmentFund",
									  "@id":"schema:InvestmentFund",
									  "layer":"pending"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialProduct",
								"description":"A financial product for the loaning of an amount of money under agreed terms and charges.",
								"name":"LoanOrCredit",
								"@id":"schema:LoanOrCredit",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:LoanOrCredit",
									  "description":"A loan in which property or real estate is used as collateral...",
									  "name":"MortgageLoan",
									  "@id":"schema:MortgageLoan",
									  "layer":"pending"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialProduct",
								"description":"A Service to transfer funds from a person or organization to a beneficiary person or organization.",
								"name":"PaymentService",
								"@id":"schema:PaymentService",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A food service, like breakfast, lunch, or dinner.",
						  "name":"FoodService",
						  "@id":"schema:FoodService",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A service provided by a government organization, e.g. food stamps, veterans benefits, etc.",
						  "name":"GovernmentService",
						  "@id":"schema:GovernmentService",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A taxi.",
						  "name":"Taxi",
						  "@id":"schema:Taxi",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"A service for a vehicle for hire with a driver for local travel...",
						  "name":"TaxiService",
						  "@id":"schema:TaxiService",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Service",
						  "description":"An application programming interface accessible over Web/Internet technologies.",
						  "name":"WebAPI",
						  "@id":"schema:WebAPI",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A means for accessing a service, e.g. a government office location, web site, or phone number.",
					"name":"ServiceChannel",
					"@id":"schema:ServiceChannel",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A SpeakableSpecification indicates (typically via xpath or cssSelector) sections of a document that are highlighted as particularly speakable...",
					"name":"SpeakableSpecification",
					"@id":"schema:SpeakableSpecification",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A StatisticalPopulation is a set of instances of a certain given type that satisfy some set of constraints...",
					"name":"StatisticalPopulation",
					"@id":"schema:StatisticalPopulation",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Structured values are used when the value of a property has a more complex structure than simply being a textual value or a reference to another thing.",
					"name":"StructuredValue",
					"@id":"schema:StructuredValue",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A CDCPMDRecord is a data structure representing a record in a CDC tabular data format\n used for hospital data reporting...",
						  "name":"CDCPMDRecord",
						  "@id":"schema:CDCPMDRecord",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A contact point—for example, a Customer Complaints department.",
						  "name":"ContactPoint",
						  "@id":"schema:ContactPoint",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:ContactPoint",
								"description":"The mailing address.",
								"name":"PostalAddress",
								"@id":"schema:PostalAddress",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A DatedMoneySpecification represents monetary values with optional start and end dates...",
						  "name":"DatedMoneySpecification",
						  "@id":"schema:DatedMoneySpecification",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"Information about the engine of the vehicle. A vehicle can have multiple engines represented by multiple engine specification entities.",
						  "name":"EngineSpecification",
						  "@id":"schema:EngineSpecification",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value representing exchange rate.",
						  "name":"ExchangeRateSpecification",
						  "@id":"schema:ExchangeRateSpecification",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"The geographic coordinates of a place or event.",
						  "name":"GeoCoordinates",
						  "@id":"schema:GeoCoordinates",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"The geographic shape of a place. A GeoShape can be described using several properties whose values are based on latitude/longitude pairs...",
						  "name":"GeoShape",
						  "@id":"schema:GeoShape",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GeoShape",
								"description":"A GeoCircle is a GeoShape representing a circular geographic area...",
								"name":"GeoCircle",
								"@id":"schema:GeoCircle",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A summary of how users have interacted with this CreativeWork...",
						  "name":"InteractionCounter",
						  "@id":"schema:InteractionCounter",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A monetary value or range. This type can be used to describe an amount of money such as $50 USD, or a range as in describing a bank account being suitable for a balance between \u00a31,000 and \u00a31,000,000 GBP, or the value of a salary, etc...",
						  "name":"MonetaryAmount",
						  "@id":"schema:MonetaryAmount",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"Nutritional information about the recipe.",
						  "name":"NutritionInformation",
						  "@id":"schema:NutritionInformation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value providing information about the opening hours of a place or a certain service inside a place...",
						  "name":"OpeningHoursSpecification",
						  "@id":"schema:OpeningHoursSpecification",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value providing information about when a certain organization or person owned a certain product.",
						  "name":"OwnershipInfo",
						  "@id":"schema:OwnershipInfo",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value representing a price or price range. Typically, only the subclasses of this type are used for markup...",
						  "name":"PriceSpecification",
						  "@id":"schema:PriceSpecification",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PriceSpecification",
								"description":"A compound price specification is one that bundles multiple prices that all apply in combination for different dimensions of consumption...",
								"name":"CompoundPriceSpecification",
								"@id":"schema:CompoundPriceSpecification",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PriceSpecification",
								"description":"The price for the delivery of an offer using a particular delivery method.",
								"name":"DeliveryChargeSpecification",
								"@id":"schema:DeliveryChargeSpecification",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PriceSpecification",
								"description":"The costs of settling the payment using a particular payment method.",
								"name":"PaymentChargeSpecification",
								"@id":"schema:PaymentChargeSpecification",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PriceSpecification",
								"description":"The price asked for a given offer by the respective organization or person.",
								"name":"UnitPriceSpecification",
								"@id":"schema:UnitPriceSpecification",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A property-value pair, e.g. representing a feature of a product or place...",
						  "name":"PropertyValue",
						  "@id":"schema:PropertyValue",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PropertyValue",
								"description":"Specifies a location feature by providing a structured value representing a feature of an accommodation as a property-value pair of varying degrees of formality.",
								"name":"LocationFeatureSpecification",
								"@id":"schema:LocationFeatureSpecification",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A point value or interval for product characteristics and other purposes.",
						  "name":"QuantitativeValue",
						  "@id":"schema:QuantitativeValue",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A statistical distribution of values.",
						  "name":"QuantitativeValueDistribution",
						  "@id":"schema:QuantitativeValueDistribution",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:QuantitativeValueDistribution",
								"description":"A statistical distribution of monetary amounts.",
								"name":"MonetaryAmountDistribution",
								"@id":"schema:MonetaryAmountDistribution",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value representing repayment.",
						  "name":"RepaymentSpecification",
						  "@id":"schema:RepaymentSpecification",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value indicating the quantity, unit of measurement, and business function of goods included in a bundle offer.",
						  "name":"TypeAndQuantityNode",
						  "@id":"schema:TypeAndQuantityNode",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:StructuredValue",
						  "description":"A structured value representing the duration and scope of services that will be provided to a customer free of charge in case of a defect or malfunction of a product.",
						  "name":"WarrantyPromise",
						  "@id":"schema:WarrantyPromise",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"Used to describe a ticket to an event, a flight, a bus ride, etc.",
					"name":"Ticket",
					"@id":"schema:Ticket",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"A trip or journey. An itinerary of visits to one or more places.",
					"name":"Trip",
					"@id":"schema:Trip",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Trip",
						  "description":"A trip on a commercial bus line.",
						  "name":"BusTrip",
						  "@id":"schema:BusTrip",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Trip",
						  "description":"An airline flight.",
						  "name":"Flight",
						  "@id":"schema:Flight",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Trip",
						  "description":"A tourist trip. A created itinerary of visits to one or more places of interest (TouristAttraction/TouristDestination) often linked by a similar theme, geographic area, or interest to a particular touristType...",
						  "name":"TouristTrip",
						  "@id":"schema:TouristTrip",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Trip",
						  "description":"A trip on a commercial train line.",
						  "name":"TrainTrip",
						  "@id":"schema:TrainTrip",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Intangible",
					"description":"An online or virtual location for attending events. For example, one may attend an online seminar or educational event...",
					"name":"VirtualLocation",
					"@id":"schema:VirtualLocation",
					"layer":"pending"
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"The most generic type of entity related to health and the practice of medicine.",
			  "name":"MedicalEntity",
			  "@id":"schema:MedicalEntity",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any part of the human body, typically a component of an anatomical system...",
					"name":"AnatomicalStructure",
					"@id":"schema:AnatomicalStructure",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"Rigid connective tissue that comprises up the skeletal structure of the human body.",
						  "name":"Bone",
						  "@id":"schema:Bone",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"Any anatomical structure which pertains to the soft nervous tissue functioning as the coordinating center of sensation and intellectual and nervous activity.",
						  "name":"BrainStructure",
						  "@id":"schema:BrainStructure",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"The anatomical location at which two or more bones make contact.",
						  "name":"Joint",
						  "@id":"schema:Joint",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"A short band of tough, flexible, fibrous connective tissue that functions to connect multiple bones, cartilages, and structurally support joints.",
						  "name":"Ligament",
						  "@id":"schema:Ligament",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"A muscle is an anatomical structure consisting of a contractile form of tissue that animals use to effect movement.",
						  "name":"Muscle",
						  "@id":"schema:Muscle",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"A common pathway for the electrochemical nerve impulses that are transmitted along each of the axons.",
						  "name":"Nerve",
						  "@id":"schema:Nerve",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AnatomicalStructure",
						  "description":"A component of the human body circulatory system comprised of an intricate network of hollow tubes that transport blood throughout the entire body.",
						  "name":"Vessel",
						  "@id":"schema:Vessel",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Vessel",
								"description":"A type of blood vessel that specifically carries blood away from the heart.",
								"name":"Artery",
								"@id":"schema:Artery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Vessel",
								"description":"A type of blood vessel that specifically carries lymph fluid unidirectionally toward the heart.",
								"name":"LymphaticVessel",
								"@id":"schema:LymphaticVessel",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Vessel",
								"description":"A type of blood vessel that specifically carries blood to the heart.",
								"name":"Vein",
								"@id":"schema:Vein",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"An anatomical system is a group of anatomical structures that work together to perform a certain task...",
					"name":"AnatomicalSystem",
					"@id":"schema:AnatomicalSystem",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A process of care involving exercise, changes to diet, fitness routines, and other lifestyle changes aimed at improving a health condition.",
					"name":"LifestyleModification",
					"@id":"schema:LifestyleModification",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LifestyleModification",
						  "description":"Any bodily activity that enhances or maintains physical fitness and overall health and wellness...",
						  "name":"PhysicalActivity",
						  "@id":"schema:PhysicalActivity",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"The causative agent(s) that are responsible for the pathophysiologic process that eventually results in a medical condition, symptom or sign...",
					"name":"MedicalCause",
					"@id":"schema:MedicalCause",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any condition of the human body that affects the normal functioning of a person, whether physically or mentally...",
					"name":"MedicalCondition",
					"@id":"schema:MedicalCondition",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalCondition",
						  "description":"An infectious disease is a clinically evident human disease resulting from the presence of pathogenic microbial agents, like pathogenic viruses, pathogenic bacteria, fungi, protozoa, multicellular parasites, and prions...",
						  "name":"InfectiousDisease",
						  "@id":"schema:InfectiousDisease",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalCondition",
						  "description":"Any feature associated or not with a medical condition. In medicine a symptom is generally subjective while a sign is objective.",
						  "name":"MedicalSignOrSymptom",
						  "@id":"schema:MedicalSignOrSymptom",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalSignOrSymptom",
								"description":"Any physical manifestation of a person's medical condition discoverable by objective diagnostic tests or physical examination.",
								"name":"MedicalSign",
								"@id":"schema:MedicalSign",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalSign",
									  "description":"Vital signs are measures of various physiological functions in order to assess the most basic body functions.",
									  "name":"VitalSign",
									  "@id":"schema:VitalSign",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalSignOrSymptom",
								"description":"Any complaint sensed and expressed by the patient (therefore defined as subjective) like stomachache, lower-back pain, or fatigue.",
								"name":"MedicalSymptom",
								"@id":"schema:MedicalSymptom",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A condition or factor that serves as a reason to withhold a certain medical therapy...",
					"name":"MedicalContraindication",
					"@id":"schema:MedicalContraindication",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any object used in a medical capacity, such as to diagnose or treat a patient.",
					"name":"MedicalDevice",
					"@id":"schema:MedicalDevice",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any recommendation made by a standard society (e.g. ACC/AHA) or consensus statement that denotes how to diagnose and treat a particular condition...",
					"name":"MedicalGuideline",
					"@id":"schema:MedicalGuideline",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalGuideline",
						  "description":"A guideline contraindication that designates a process as harmful and where quality of the data supporting the contraindication is sound.",
						  "name":"MedicalGuidelineContraindication",
						  "@id":"schema:MedicalGuidelineContraindication",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalGuideline",
						  "description":"A guideline recommendation that is regarded as efficacious and where quality of the data supporting the recommendation is sound.",
						  "name":"MedicalGuidelineRecommendation",
						  "@id":"schema:MedicalGuidelineRecommendation",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A condition or factor that indicates use of a medical therapy, including signs, symptoms, risk factors, anatomical states, etc.",
					"name":"MedicalIndication",
					"@id":"schema:MedicalIndication",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIndication",
						  "description":"An indication for a medical therapy that has been formally specified or approved by a regulatory body that regulates use of the therapy; for example, the US FDA approves indications for most drugs in the US.",
						  "name":"ApprovedIndication",
						  "@id":"schema:ApprovedIndication",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIndication",
						  "description":"An indication for preventing an underlying condition, symptom, etc.",
						  "name":"PreventionIndication",
						  "@id":"schema:PreventionIndication",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIndication",
						  "description":"An indication for treating an underlying condition, symptom, etc.",
						  "name":"TreatmentIndication",
						  "@id":"schema:TreatmentIndication",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A utility class that serves as the umbrella for a number of 'intangible' things in the medical space.",
					"name":"MedicalIntangible",
					"@id":"schema:MedicalIntangible",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIntangible",
						  "description":"An alternative, closely-related condition typically considered later in the differential diagnosis process along with the signs that are used to distinguish it.",
						  "name":"DDxElement",
						  "@id":"schema:DDxElement",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIntangible",
						  "description":"A specific dosing schedule for a drug or supplement.",
						  "name":"DoseSchedule",
						  "@id":"schema:DoseSchedule",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DoseSchedule",
								"description":"The maximum dosing schedule considered safe for a drug or supplement as recommended by an authority or by the drug/supplement's manufacturer...",
								"name":"MaximumDoseSchedule",
								"@id":"schema:MaximumDoseSchedule",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DoseSchedule",
								"description":"A recommended dosing schedule for a drug or supplement as prescribed or recommended by an authority or by the drug/supplement's manufacturer...",
								"name":"RecommendedDoseSchedule",
								"@id":"schema:RecommendedDoseSchedule",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:DoseSchedule",
								"description":"A patient-reported or observed dosing schedule for a drug or supplement.",
								"name":"ReportedDoseSchedule",
								"@id":"schema:ReportedDoseSchedule",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIntangible",
						  "description":"The legal availability status of a medical drug.",
						  "name":"DrugLegalStatus",
						  "@id":"schema:DrugLegalStatus",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIntangible",
						  "description":"A specific strength in which a medical drug is available in a specific country.",
						  "name":"DrugStrength",
						  "@id":"schema:DrugStrength",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalIntangible",
						  "description":"A stage of a medical condition, such as 'Stage IIIa'.",
						  "name":"MedicalConditionStage",
						  "@id":"schema:MedicalConditionStage",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A process of care used in either a diagnostic, therapeutic, preventive or palliative capacity that relies on invasive (surgical), non-invasive, or other techniques.",
					"name":"MedicalProcedure",
					"@id":"schema:MedicalProcedure",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalProcedure",
						  "description":"A medical procedure intended primarily for diagnostic, as opposed to therapeutic, purposes.",
						  "name":"DiagnosticProcedure",
						  "@id":"schema:DiagnosticProcedure",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalProcedure",
						  "description":"A medical procedure intended primarily for palliative purposes, aimed at relieving the symptoms of an underlying health condition.",
						  "name":"PalliativeProcedure",
						  "@id":"schema:PalliativeProcedure",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalProcedure",
						  "description":"A medical procedure involving an incision with instruments; performed for diagnose, or therapeutic purposes.",
						  "name":"SurgicalProcedure",
						  "@id":"schema:SurgicalProcedure",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalProcedure",
						  "description":"A medical procedure intended primarily for therapeutic purposes, aimed at improving a health condition.",
						  "name":"TherapeuticProcedure",
						  "@id":"schema:TherapeuticProcedure",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:TherapeuticProcedure",
								"description":"Any medical intervention designed to prevent, treat, and cure human diseases and medical conditions, including both curative and palliative therapies...",
								"name":"MedicalTherapy",
								"@id":"schema:MedicalTherapy",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTherapy",
									  "description":"A treatment of people with physical, emotional, or social problems, using purposeful activity to help them overcome or learn to deal with their problems.",
									  "name":"OccupationalTherapy",
									  "@id":"schema:OccupationalTherapy",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTherapy",
									  "description":"A process of progressive physical care and rehabilitation aimed at improving a health condition.",
									  "name":"PhysicalTherapy",
									  "@id":"schema:PhysicalTherapy",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTherapy",
									  "description":"A process of care using radiation aimed at improving a health condition.",
									  "name":"RadiationTherapy",
									  "@id":"schema:RadiationTherapy",
									  "layer":"core"
								   },
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalTherapy",
									  "description":"The therapy that is concerned with the maintenance or improvement of respiratory function (as in patients with pulmonary disease).",
									  "name":"RespiratoryTherapy",
									  "@id":"schema:RespiratoryTherapy",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:TherapeuticProcedure",
								"description":"A process of care relying upon counseling, dialogue and communication aimed at improving a mental health condition without use of drugs.",
								"name":"PsychologicalTreatment",
								"@id":"schema:PsychologicalTreatment",
								"layer":"core"
							 }
						  ]
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any rule set or interactive tool for estimating the risk of developing a complication or condition.",
					"name":"MedicalRiskEstimator",
					"@id":"schema:MedicalRiskEstimator",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalRiskEstimator",
						  "description":"A complex mathematical calculation requiring an online calculator, used to assess prognosis...",
						  "name":"MedicalRiskCalculator",
						  "@id":"schema:MedicalRiskCalculator",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalRiskEstimator",
						  "description":"A simple system that adds up the number of risk factors to yield a score that is associated with prognosis, e...",
						  "name":"MedicalRiskScore",
						  "@id":"schema:MedicalRiskScore",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A risk factor is anything that increases a person's likelihood of developing or contracting a disease, medical condition, or complication.",
					"name":"MedicalRiskFactor",
					"@id":"schema:MedicalRiskFactor",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"A medical study is an umbrella type covering all kinds of research studies relating to human medicine or health, including observational studies and interventional trials and registries, randomized, controlled or not...",
					"name":"MedicalStudy",
					"@id":"schema:MedicalStudy",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalStudy",
						  "description":"An observational study is a type of medical study that attempts to infer the possible effect of a treatment through observation of a cohort of subjects over a period of time...",
						  "name":"MedicalObservationalStudy",
						  "@id":"schema:MedicalObservationalStudy",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalStudy",
						  "description":"A medical trial is a type of medical study that uses scientific process used to compare the safety and efficacy of medical therapies or medical procedures...",
						  "name":"MedicalTrial",
						  "@id":"schema:MedicalTrial",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any medical test, typically performed for diagnostic purposes.",
					"name":"MedicalTest",
					"@id":"schema:MedicalTest",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalTest",
						  "description":"A medical test performed on a sample of a patient's blood.",
						  "name":"BloodTest",
						  "@id":"schema:BloodTest",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalTest",
						  "description":"Any medical imaging modality typically used for diagnostic purposes.",
						  "name":"ImagingTest",
						  "@id":"schema:ImagingTest",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalTest",
						  "description":"Any collection of tests commonly ordered together.",
						  "name":"MedicalTestPanel",
						  "@id":"schema:MedicalTestPanel",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalTest",
						  "description":"A medical test performed by a laboratory that typically involves examination of a tissue sample by a pathologist.",
						  "name":"PathologyTest",
						  "@id":"schema:PathologyTest",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Any matter of defined composition that has discrete existence, whose origin may be biological, mineral or chemical.",
					"name":"Substance",
					"@id":"schema:Substance",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Substance",
						  "description":"A product taken by mouth that contains a dietary ingredient intended to supplement the diet...",
						  "name":"DietarySupplement",
						  "@id":"schema:DietarySupplement",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Substance",
						  "description":"A chemical or biologic substance, used as a medical therapy, that has a physiological effect on an organism...",
						  "name":"Drug",
						  "@id":"schema:Drug",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:MedicalEntity",
					"description":"Anatomical features that can be observed by sight (without dissection), including the form and proportions of the human body as well as surface landmarks that correspond to deeper subcutaneous structures...",
					"name":"SuperficialAnatomy",
					"@id":"schema:SuperficialAnatomy",
					"layer":"core"
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"An organization such as a school, NGO, corporation, club, etc.",
			  "name":"Organization",
			  "@id":"schema:Organization",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"An organization that provides flights for passengers.",
					"name":"Airline",
					"@id":"schema:Airline",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A Consortium is a membership Organization whose members are typically Organizations.",
					"name":"Consortium",
					"@id":"schema:Consortium",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"Organization: A business corporation.",
					"name":"Corporation",
					"@id":"schema:Corporation",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"An educational organization.",
					"name":"EducationalOrganization",
					"@id":"schema:EducationalOrganization",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOrganization",
						  "description":"A college, university, or other third-level educational institution.",
						  "name":"CollegeOrUniversity",
						  "@id":"schema:CollegeOrUniversity",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOrganization",
						  "description":"An elementary school.",
						  "name":"ElementarySchool",
						  "@id":"schema:ElementarySchool",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOrganization",
						  "description":"A high school.",
						  "name":"HighSchool",
						  "@id":"schema:HighSchool",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOrganization",
						  "description":"A middle school (typically for children aged around 11-14, although this varies somewhat).",
						  "name":"MiddleSchool",
						  "@id":"schema:MiddleSchool",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOrganization",
						  "description":"A preschool.",
						  "name":"Preschool",
						  "@id":"schema:Preschool",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:EducationalOrganization",
						  "description":"A school.",
						  "name":"School",
						  "@id":"schema:School",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A FundingScheme combines organizational, project and policy aspects of grant-based funding\n that sets guidelines, principles and mechanisms to support other kinds of projects and activities...",
					"name":"FundingScheme",
					"@id":"schema:FundingScheme",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A governmental organization or agency.",
					"name":"GovernmentOrganization",
					"@id":"schema:GovernmentOrganization",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A LibrarySystem is a collaborative system amongst several libraries.",
					"name":"LibrarySystem",
					"@id":"schema:LibrarySystem",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A particular physical business or branch of an organization. Examples of LocalBusiness include a restaurant, a particular branch of a restaurant chain, a branch of a bank, a medical practice, a club, a bowling alley, etc.",
					"name":"LocalBusiness",
					"@id":"schema:LocalBusiness",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"Animal shelter.",
						  "name":"AnimalShelter",
						  "@id":"schema:AnimalShelter",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"An organization with archival holdings. An organization which keeps and preserves archival material and typically makes it accessible to the public.",
						  "name":"ArchiveOrganization",
						  "@id":"schema:ArchiveOrganization",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"Car repair, sales, or parts.",
						  "name":"AutomotiveBusiness",
						  "@id":"schema:AutomotiveBusiness",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"Auto body shop.",
								"name":"AutoBodyShop",
								"@id":"schema:AutoBodyShop",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"An car dealership.",
								"name":"AutoDealer",
								"@id":"schema:AutoDealer",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"An auto parts store.",
								"name":"AutoPartsStore",
								"@id":"schema:AutoPartsStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"A car rental business.",
								"name":"AutoRental",
								"@id":"schema:AutoRental",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"Car repair business.",
								"name":"AutoRepair",
								"@id":"schema:AutoRepair",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"A car wash business.",
								"name":"AutoWash",
								"@id":"schema:AutoWash",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"A gas station.",
								"name":"GasStation",
								"@id":"schema:GasStation",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"A motorcycle dealer.",
								"name":"MotorcycleDealer",
								"@id":"schema:MotorcycleDealer",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:AutomotiveBusiness",
								"description":"A motorcycle repair shop.",
								"name":"MotorcycleRepair",
								"@id":"schema:MotorcycleRepair",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A Childcare center.",
						  "name":"ChildCare",
						  "@id":"schema:ChildCare",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A dentist.",
						  "name":"Dentist",
						  "@id":"schema:Dentist",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A dry-cleaning business.",
						  "name":"DryCleaningOrLaundry",
						  "@id":"schema:DryCleaningOrLaundry",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"An emergency service, such as a fire station or ER.",
						  "name":"EmergencyService",
						  "@id":"schema:EmergencyService",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EmergencyService",
								"description":"A fire station. With firemen.",
								"name":"FireStation",
								"@id":"schema:FireStation",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EmergencyService",
								"description":"A hospital.",
								"name":"Hospital",
								"@id":"schema:Hospital",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EmergencyService",
								"description":"A police station.",
								"name":"PoliceStation",
								"@id":"schema:PoliceStation",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"An employment agency.",
						  "name":"EmploymentAgency",
						  "@id":"schema:EmploymentAgency",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A business providing entertainment.",
						  "name":"EntertainmentBusiness",
						  "@id":"schema:EntertainmentBusiness",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"An adult entertainment establishment.",
								"name":"AdultEntertainment",
								"@id":"schema:AdultEntertainment",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"An amusement park.",
								"name":"AmusementPark",
								"@id":"schema:AmusementPark",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"An art gallery.",
								"name":"ArtGallery",
								"@id":"schema:ArtGallery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"A casino.",
								"name":"Casino",
								"@id":"schema:Casino",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"A comedy club.",
								"name":"ComedyClub",
								"@id":"schema:ComedyClub",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"A movie theater.",
								"name":"MovieTheater",
								"@id":"schema:MovieTheater",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:EntertainmentBusiness",
								"description":"A nightclub or discotheque.",
								"name":"NightClub",
								"@id":"schema:NightClub",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"Financial services business.",
						  "name":"FinancialService",
						  "@id":"schema:FinancialService",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialService",
								"description":"Accountancy business.\n\nAs a LocalBusiness it can be described as a provider of one or more Service(s).",
								"name":"AccountingService",
								"@id":"schema:AccountingService",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialService",
								"description":"ATM/cash machine.",
								"name":"AutomatedTeller",
								"@id":"schema:AutomatedTeller",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialService",
								"description":"Bank or credit union.",
								"name":"BankOrCreditUnion",
								"@id":"schema:BankOrCreditUnion",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FinancialService",
								"description":"An Insurance agency.",
								"name":"InsuranceAgency",
								"@id":"schema:InsuranceAgency",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A food-related business.",
						  "name":"FoodEstablishment",
						  "@id":"schema:FoodEstablishment",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A bakery.",
								"name":"Bakery",
								"@id":"schema:Bakery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A bar or pub.",
								"name":"BarOrPub",
								"@id":"schema:BarOrPub",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"Brewery.",
								"name":"Brewery",
								"@id":"schema:Brewery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A cafe or coffee shop.",
								"name":"CafeOrCoffeeShop",
								"@id":"schema:CafeOrCoffeeShop",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A distillery.",
								"name":"Distillery",
								"@id":"schema:Distillery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A fast-food restaurant.",
								"name":"FastFoodRestaurant",
								"@id":"schema:FastFoodRestaurant",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"An ice cream shop.",
								"name":"IceCreamShop",
								"@id":"schema:IceCreamShop",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A restaurant.",
								"name":"Restaurant",
								"@id":"schema:Restaurant",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:FoodEstablishment",
								"description":"A winery.",
								"name":"Winery",
								"@id":"schema:Winery",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A government office—for example, an IRS or DMV office.",
						  "name":"GovernmentOffice",
						  "@id":"schema:GovernmentOffice",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GovernmentOffice",
								"description":"A post office.",
								"name":"PostOffice",
								"@id":"schema:PostOffice",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"Health and beauty.",
						  "name":"HealthAndBeautyBusiness",
						  "@id":"schema:HealthAndBeautyBusiness",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAndBeautyBusiness",
								"description":"Beauty salon.",
								"name":"BeautySalon",
								"@id":"schema:BeautySalon",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAndBeautyBusiness",
								"description":"A day spa.",
								"name":"DaySpa",
								"@id":"schema:DaySpa",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAndBeautyBusiness",
								"description":"A hair salon.",
								"name":"HairSalon",
								"@id":"schema:HairSalon",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAndBeautyBusiness",
								"description":"A health club.",
								"name":"HealthClub",
								"@id":"schema:HealthClub",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAndBeautyBusiness",
								"description":"A nail salon.",
								"name":"NailSalon",
								"@id":"schema:NailSalon",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HealthAndBeautyBusiness",
								"description":"A tattoo parlor.",
								"name":"TattooParlor",
								"@id":"schema:TattooParlor",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A construction business.\n\nA HomeAndConstructionBusiness is a LocalBusiness that provides services around homes and buildings...",
						  "name":"HomeAndConstructionBusiness",
						  "@id":"schema:HomeAndConstructionBusiness",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"An electrician.",
								"name":"Electrician",
								"@id":"schema:Electrician",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A general contractor.",
								"name":"GeneralContractor",
								"@id":"schema:GeneralContractor",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A business that provide Heating, Ventilation and Air Conditioning services.",
								"name":"HVACBusiness",
								"@id":"schema:HVACBusiness",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A house painting service.",
								"name":"HousePainter",
								"@id":"schema:HousePainter",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A locksmith.",
								"name":"Locksmith",
								"@id":"schema:Locksmith",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A moving company.",
								"name":"MovingCompany",
								"@id":"schema:MovingCompany",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A plumbing service.",
								"name":"Plumber",
								"@id":"schema:Plumber",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:HomeAndConstructionBusiness",
								"description":"A roofing contractor.",
								"name":"RoofingContractor",
								"@id":"schema:RoofingContractor",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"An internet cafe.",
						  "name":"InternetCafe",
						  "@id":"schema:InternetCafe",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A LegalService is a business that provides legally-oriented services, advice and representation, e...",
						  "name":"LegalService",
						  "@id":"schema:LegalService",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalService",
								"description":"Professional service: Attorney. \n\nThis type is deprecated - LegalService is more inclusive and less ambiguous.",
								"name":"Attorney",
								"@id":"schema:Attorney",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LegalService",
								"description":"A notary.",
								"name":"Notary",
								"@id":"schema:Notary",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A library.",
						  "name":"Library",
						  "@id":"schema:Library",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A lodging business, such as a motel, hotel, or inn.",
						  "name":"LodgingBusiness",
						  "@id":"schema:LodgingBusiness",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LodgingBusiness",
								"description":"Bed and breakfast.\n\nSee also the dedicated document on the use of schema...",
								"name":"BedAndBreakfast",
								"@id":"schema:BedAndBreakfast",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LodgingBusiness",
								"description":"A camping site, campsite, or Campground is a place used for overnight stay in the outdoors, typically containing individual CampingPitch locations...",
								"name":"Campground",
								"@id":"schema:Campground",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LodgingBusiness",
								"description":"A hostel - cheap accommodation, often in shared dormitories.\n\nSee also the dedicated document on the use of schema...",
								"name":"Hostel",
								"@id":"schema:Hostel",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LodgingBusiness",
								"description":"A hotel is an establishment that provides lodging paid on a short-term basis (Source: Wikipedia, the free encyclopedia, see http://en...",
								"name":"Hotel",
								"@id":"schema:Hotel",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LodgingBusiness",
								"description":"A motel.\n\nSee also the dedicated document on the use of schema...",
								"name":"Motel",
								"@id":"schema:Motel",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:LodgingBusiness",
								"description":"A resort is a place used for relaxation or recreation, attracting visitors for holidays or vacations...",
								"name":"Resort",
								"@id":"schema:Resort",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A particular physical or virtual business of an organization for medical purposes...",
						  "name":"MedicalBusiness",
						  "@id":"schema:MedicalBusiness",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A field of public health focusing on improving health characteristics of a defined population in relation with their geographical or environment areas",
								"name":"CommunityHealth",
								"@id":"schema:CommunityHealth",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that pertains to diagnosis and treatment of disorders of skin.",
								"name":"Dermatology",
								"@id":"schema:Dermatology",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"Dietetic and nutrition as a medical speciality.",
								"name":"DietNutrition",
								"@id":"schema:DietNutrition",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that deals with the evaluation and initial treatment of medical conditions caused by trauma or sudden illness.",
								"name":"Emergency",
								"@id":"schema:Emergency",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that is concerned with the diagnosis and treatment of diseases, debilities and provision of care to the aged.",
								"name":"Geriatric",
								"@id":"schema:Geriatric",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that pertains to the health care of women, particularly in the diagnosis and treatment of disorders affecting the female reproductive system.",
								"name":"Gynecologic",
								"@id":"schema:Gynecologic",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A facility, often associated with a hospital or medical school, that is devoted to the specific diagnosis and/or healthcare...",
								"name":"MedicalClinic",
								"@id":"schema:MedicalClinic",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:MedicalClinic",
									  "description":"A CovidTestingFacility is a MedicalClinic where testing for the COVID-19 Coronavirus\n disease is available...",
									  "name":"CovidTestingFacility",
									  "@id":"schema:CovidTestingFacility",
									  "layer":"pending"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A nurse-like health profession that deals with pregnancy, childbirth, and the postpartum period (including care of the newborn), besides sexual and reproductive health of women throughout their lives.",
								"name":"Midwifery",
								"@id":"schema:Midwifery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A health profession of a person formally educated and trained in the care of the sick or infirm person.",
								"name":"Nursing",
								"@id":"schema:Nursing",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that specializes in the care of women during the prenatal and postnatal care and with the delivery of the child.",
								"name":"Obstetric",
								"@id":"schema:Obstetric",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that deals with benign and malignant tumors, including the study of their development, diagnosis, treatment and prevention.",
								"name":"Oncologic",
								"@id":"schema:Oncologic",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A store that sells reading glasses and similar devices for improving vision.",
								"name":"Optician",
								"@id":"schema:Optician",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"The science or practice of testing visual acuity and prescribing corrective lenses.",
								"name":"Optometric",
								"@id":"schema:Optometric",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that is concerned with the ear, nose and throat and their respective disease states.",
								"name":"Otolaryngologic",
								"@id":"schema:Otolaryngologic",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that specializes in the care of infants, children and adolescents.",
								"name":"Pediatric",
								"@id":"schema:Pediatric",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A pharmacy or drugstore.",
								"name":"Pharmacy",
								"@id":"schema:Pharmacy",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A doctor's office.",
								"name":"Physician",
								"@id":"schema:Physician",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"The practice of treatment of disease, injury, or deformity by physical methods such as massage, heat treatment, and exercise rather than by drugs or surgery.",
								"name":"Physiotherapy",
								"@id":"schema:Physiotherapy",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that pertains to therapeutic or cosmetic repair or re-formation of missing, injured or malformed tissues or body parts by manual and instrumental means.",
								"name":"PlasticSurgery",
								"@id":"schema:PlasticSurgery",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"Podiatry is the care of the human foot, especially the diagnosis and treatment of foot disorders.",
								"name":"Podiatric",
								"@id":"schema:Podiatric",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"The medical care by a physician, or other health-care professional, who is the patient's first contact with the health-care system and who may recommend a specialist if necessary.",
								"name":"PrimaryCare",
								"@id":"schema:PrimaryCare",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"A specific branch of medical science that is concerned with the study, treatment, and prevention of mental illness, using both medical and psychological therapies.",
								"name":"Psychiatric",
								"@id":"schema:Psychiatric",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:MedicalBusiness",
								"description":"Branch of medicine that pertains to the health services to improve and protect community health, especially epidemiology, sanitation, immunization, and preventive medicine.",
								"name":"PublicHealth",
								"@id":"schema:PublicHealth",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"Original definition: \"provider of professional services.\"\n\nThe general ProfessionalService type for local businesses was deprecated due to confusion with Service...",
						  "name":"ProfessionalService",
						  "@id":"schema:ProfessionalService",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A radio station.",
						  "name":"RadioStation",
						  "@id":"schema:RadioStation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A real-estate agent.",
						  "name":"RealEstateAgent",
						  "@id":"schema:RealEstateAgent",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A recycling center.",
						  "name":"RecyclingCenter",
						  "@id":"schema:RecyclingCenter",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A self-storage facility.",
						  "name":"SelfStorage",
						  "@id":"schema:SelfStorage",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A shopping center or mall.",
						  "name":"ShoppingCenter",
						  "@id":"schema:ShoppingCenter",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A sports location, such as a playing field.",
						  "name":"SportsActivityLocation",
						  "@id":"schema:SportsActivityLocation",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A bowling alley.",
								"name":"BowlingAlley",
								"@id":"schema:BowlingAlley",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A gym.",
								"name":"ExerciseGym",
								"@id":"schema:ExerciseGym",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A golf course.",
								"name":"GolfCourse",
								"@id":"schema:GolfCourse",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A public swimming pool.",
								"name":"PublicSwimmingPool",
								"@id":"schema:PublicSwimmingPool",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A ski resort.",
								"name":"SkiResort",
								"@id":"schema:SkiResort",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A sports club.",
								"name":"SportsClub",
								"@id":"schema:SportsClub",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A stadium.",
								"name":"StadiumOrArena",
								"@id":"schema:StadiumOrArena",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:SportsActivityLocation",
								"description":"A tennis complex.",
								"name":"TennisComplex",
								"@id":"schema:TennisComplex",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A retail good store.",
						  "name":"Store",
						  "@id":"schema:Store",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A bike store.",
								"name":"BikeStore",
								"@id":"schema:BikeStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A bookstore.",
								"name":"BookStore",
								"@id":"schema:BookStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A clothing store.",
								"name":"ClothingStore",
								"@id":"schema:ClothingStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A computer store.",
								"name":"ComputerStore",
								"@id":"schema:ComputerStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A convenience store.",
								"name":"ConvenienceStore",
								"@id":"schema:ConvenienceStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A department store.",
								"name":"DepartmentStore",
								"@id":"schema:DepartmentStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"An electronics store.",
								"name":"ElectronicsStore",
								"@id":"schema:ElectronicsStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A florist.",
								"name":"Florist",
								"@id":"schema:Florist",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A furniture store.",
								"name":"FurnitureStore",
								"@id":"schema:FurnitureStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A garden store.",
								"name":"GardenStore",
								"@id":"schema:GardenStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A grocery store.",
								"name":"GroceryStore",
								"@id":"schema:GroceryStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A hardware store.",
								"name":"HardwareStore",
								"@id":"schema:HardwareStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A store that sells materials useful or necessary for various hobbies.",
								"name":"HobbyShop",
								"@id":"schema:HobbyShop",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A home goods store.",
								"name":"HomeGoodsStore",
								"@id":"schema:HomeGoodsStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A jewelry store.",
								"name":"JewelryStore",
								"@id":"schema:JewelryStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A shop that sells alcoholic drinks such as wine, beer, whisky and other spirits.",
								"name":"LiquorStore",
								"@id":"schema:LiquorStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A men's clothing store.",
								"name":"MensClothingStore",
								"@id":"schema:MensClothingStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A store that sells mobile phones and related accessories.",
								"name":"MobilePhoneStore",
								"@id":"schema:MobilePhoneStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A movie rental store.",
								"name":"MovieRentalStore",
								"@id":"schema:MovieRentalStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A music store.",
								"name":"MusicStore",
								"@id":"schema:MusicStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"An office equipment store.",
								"name":"OfficeEquipmentStore",
								"@id":"schema:OfficeEquipmentStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"An outlet store.",
								"name":"OutletStore",
								"@id":"schema:OutletStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A shop that will buy, or lend money against the security of, personal possessions.",
								"name":"PawnShop",
								"@id":"schema:PawnShop",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A pet store.",
								"name":"PetStore",
								"@id":"schema:PetStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A shoe store.",
								"name":"ShoeStore",
								"@id":"schema:ShoeStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A sporting goods store.",
								"name":"SportingGoodsStore",
								"@id":"schema:SportingGoodsStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A tire shop.",
								"name":"TireShop",
								"@id":"schema:TireShop",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A toy store.",
								"name":"ToyStore",
								"@id":"schema:ToyStore",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Store",
								"description":"A wholesale store.",
								"name":"WholesaleStore",
								"@id":"schema:WholesaleStore",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A television station.",
						  "name":"TelevisionStation",
						  "@id":"schema:TelevisionStation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A tourist information center.",
						  "name":"TouristInformationCenter",
						  "@id":"schema:TouristInformationCenter",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:LocalBusiness",
						  "description":"A travel agency.",
						  "name":"TravelAgency",
						  "@id":"schema:TravelAgency",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A medical organization (physical or not), such as hospital, institution or clinic.",
					"name":"MedicalOrganization",
					"@id":"schema:MedicalOrganization",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalOrganization",
						  "description":"A medical laboratory that offers on-site or off-site diagnostic services.",
						  "name":"DiagnosticLab",
						  "@id":"schema:DiagnosticLab",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:MedicalOrganization",
						  "description":"A vet's office.",
						  "name":"VeterinaryCare",
						  "@id":"schema:VeterinaryCare",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"Organization: Non-governmental Organization.",
					"name":"NGO",
					"@id":"schema:NGO",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A News/Media organization such as a newspaper or TV station.",
					"name":"NewsMediaOrganization",
					"@id":"schema:NewsMediaOrganization",
					"layer":"pending"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A performance group, such as a band, an orchestra, or a circus.",
					"name":"PerformingGroup",
					"@id":"schema:PerformingGroup",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PerformingGroup",
						  "description":"A dance group—for example, the Alvin Ailey Dance Theater or Riverdance.",
						  "name":"DanceGroup",
						  "@id":"schema:DanceGroup",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PerformingGroup",
						  "description":"A musical group, such as a band, an orchestra, or a choir. Can also be a solo musician.",
						  "name":"MusicGroup",
						  "@id":"schema:MusicGroup",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:PerformingGroup",
						  "description":"A theater group or company, for example, the Royal Shakespeare Company or Druid Theatre.",
						  "name":"TheaterGroup",
						  "@id":"schema:TheaterGroup",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"An enterprise (potentially individual but typically collaborative), planned to achieve a particular aim...",
					"name":"Project",
					"@id":"schema:Project",
					"layer":"pending",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Project",
						  "description":"A FundingAgency is an organization that implements one or more FundingSchemes and manages\n the granting process (via Grants, typically MonetaryGrants)...",
						  "name":"FundingAgency",
						  "@id":"schema:FundingAgency",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Project",
						  "description":"A Research project.",
						  "name":"ResearchProject",
						  "@id":"schema:ResearchProject",
						  "layer":"pending"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"Represents the collection of all sports organizations, including sports teams, governing bodies, and sports associations.",
					"name":"SportsOrganization",
					"@id":"schema:SportsOrganization",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:SportsOrganization",
						  "description":"Organization: Sports team.",
						  "name":"SportsTeam",
						  "@id":"schema:SportsTeam",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Organization",
					"description":"A Workers Union (also known as a Labor Union, Labour Union, or Trade Union) is an organization that promotes the interests of its worker members by collectively bargaining with management, organizing, and political lobbying.",
					"name":"WorkersUnion",
					"@id":"schema:WorkersUnion",
					"layer":"core"
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"A person (alive, dead, undead, or fictional).",
			  "name":"Person",
			  "@id":"schema:Person",
			  "layer":"core"
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"Entities that have a somewhat fixed, physical extension.",
			  "name":"Place",
			  "@id":"schema:Place",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"An accommodation is a place that can accommodate human beings, e...",
					"name":"Accommodation",
					"@id":"schema:Accommodation",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Accommodation",
						  "description":"An apartment (in American English) or flat (in British English) is a self-contained housing unit (a type of residential real estate) that occupies only part of a building (Source: Wikipedia, the free encyclopedia, see http://en...",
						  "name":"Apartment",
						  "@id":"schema:Apartment",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Accommodation",
						  "description":"A CampingPitch is an individual place for overnight stay in the outdoors, typically being part of a larger camping site, or Campground...",
						  "name":"CampingPitch",
						  "@id":"schema:CampingPitch",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Accommodation",
						  "description":"A house is a building or structure that has the ability to be occupied for habitation by humans or other creatures (Source: Wikipedia, the free encyclopedia, see http://en...",
						  "name":"House",
						  "@id":"schema:House",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:House",
								"description":"Residence type: Single-family home.",
								"name":"SingleFamilyResidence",
								"@id":"schema:SingleFamilyResidence",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Accommodation",
						  "description":"A room is a distinguishable space within a structure, usually separated from other spaces by interior walls...",
						  "name":"Room",
						  "@id":"schema:Room",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Room",
								"description":"A hotel room is a single room in a hotel.\n\nSee also the dedicated document on the use of schema...",
								"name":"HotelRoom",
								"@id":"schema:HotelRoom",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:Room",
								"description":"A meeting room, conference room, or conference hall is a room provided for singular events such as business conferences and meetings (Source: Wikipedia, the free encyclopedia, see http://en...",
								"name":"MeetingRoom",
								"@id":"schema:MeetingRoom",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Accommodation",
						  "description":"A suite in a hotel or other public accommodation, denotes a class of luxury accommodations, the key feature of which is multiple rooms (Source: Wikipedia, the free encyclopedia, see http://en...",
						  "name":"Suite",
						  "@id":"schema:Suite",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"A geographical region, typically under the jurisdiction of a particular government.",
					"name":"AdministrativeArea",
					"@id":"schema:AdministrativeArea",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AdministrativeArea",
						  "description":"A city or town.",
						  "name":"City",
						  "@id":"schema:City",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AdministrativeArea",
						  "description":"A country.",
						  "name":"Country",
						  "@id":"schema:Country",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AdministrativeArea",
						  "description":"A School District is an administrative area for the administration of schools.",
						  "name":"SchoolDistrict",
						  "@id":"schema:SchoolDistrict",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:AdministrativeArea",
						  "description":"A state or province of a country.",
						  "name":"State",
						  "@id":"schema:State",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"A public structure, such as a town hall or concert hall.",
					"name":"CivicStructure",
					"@id":"schema:CivicStructure",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"An airport.",
						  "name":"Airport",
						  "@id":"schema:Airport",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"Aquarium.",
						  "name":"Aquarium",
						  "@id":"schema:Aquarium",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"Beach.",
						  "name":"Beach",
						  "@id":"schema:Beach",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A bridge.",
						  "name":"Bridge",
						  "@id":"schema:Bridge",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A bus station.",
						  "name":"BusStation",
						  "@id":"schema:BusStation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A bus stop.",
						  "name":"BusStop",
						  "@id":"schema:BusStop",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A graveyard.",
						  "name":"Cemetery",
						  "@id":"schema:Cemetery",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A crematorium.",
						  "name":"Crematorium",
						  "@id":"schema:Crematorium",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"An event venue.",
						  "name":"EventVenue",
						  "@id":"schema:EventVenue",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A government building.",
						  "name":"GovernmentBuilding",
						  "@id":"schema:GovernmentBuilding",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GovernmentBuilding",
								"description":"A city hall.",
								"name":"CityHall",
								"@id":"schema:CityHall",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GovernmentBuilding",
								"description":"A courthouse.",
								"name":"Courthouse",
								"@id":"schema:Courthouse",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GovernmentBuilding",
								"description":"A defence establishment, such as an army or navy base.",
								"name":"DefenceEstablishment",
								"@id":"schema:DefenceEstablishment",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GovernmentBuilding",
								"description":"An embassy.",
								"name":"Embassy",
								"@id":"schema:Embassy",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:GovernmentBuilding",
								"description":"A legislative building—for example, the state capitol.",
								"name":"LegislativeBuilding",
								"@id":"schema:LegislativeBuilding",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A museum.",
						  "name":"Museum",
						  "@id":"schema:Museum",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A music venue.",
						  "name":"MusicVenue",
						  "@id":"schema:MusicVenue",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A park.",
						  "name":"Park",
						  "@id":"schema:Park",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A parking lot or other parking facility.",
						  "name":"ParkingFacility",
						  "@id":"schema:ParkingFacility",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A theater or other performing art center.",
						  "name":"PerformingArtsTheater",
						  "@id":"schema:PerformingArtsTheater",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"Place of worship, such as a church, synagogue, or mosque.",
						  "name":"PlaceOfWorship",
						  "@id":"schema:PlaceOfWorship",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlaceOfWorship",
								"description":"A Buddhist temple.",
								"name":"BuddhistTemple",
								"@id":"schema:BuddhistTemple",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlaceOfWorship",
								"description":"A church.",
								"name":"Church",
								"@id":"schema:Church",
								"layer":"core",
								"children":[
								   {
									  "@type":"rdfs:Class",
									  "rdfs:subClassOf":"schema:Church",
									  "description":"A Catholic church.",
									  "name":"CatholicChurch",
									  "@id":"schema:CatholicChurch",
									  "layer":"core"
								   }
								]
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlaceOfWorship",
								"description":"A Hindu temple.",
								"name":"HinduTemple",
								"@id":"schema:HinduTemple",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlaceOfWorship",
								"description":"A mosque.",
								"name":"Mosque",
								"@id":"schema:Mosque",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:PlaceOfWorship",
								"description":"A synagogue.",
								"name":"Synagogue",
								"@id":"schema:Synagogue",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A playground.",
						  "name":"Playground",
						  "@id":"schema:Playground",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A public toilet is a room or small building containing one or more toilets (and possibly also urinals) which is available for use by the general public, or by customers or employees of certain businesses.",
						  "name":"PublicToilet",
						  "@id":"schema:PublicToilet",
						  "layer":"pending"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A place offering space for \"Recreational Vehicles\", Caravans, mobile homes and the like.",
						  "name":"RVPark",
						  "@id":"schema:RVPark",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A subway station.",
						  "name":"SubwayStation",
						  "@id":"schema:SubwayStation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A taxi stand.",
						  "name":"TaxiStand",
						  "@id":"schema:TaxiStand",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A train station.",
						  "name":"TrainStation",
						  "@id":"schema:TrainStation",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:CivicStructure",
						  "description":"A zoo.",
						  "name":"Zoo",
						  "@id":"schema:Zoo",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"A landform or physical feature. Landform elements include mountains, plains, lakes, rivers, seascape and oceanic waterbody interface features such as bays, peninsulas, seas and so forth, including sub-aqueous terrain features such as submersed mountain ranges, volcanoes, and the great ocean basins.",
					"name":"Landform",
					"@id":"schema:Landform",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Landform",
						  "description":"A body of water, such as a sea, ocean, or lake.",
						  "name":"BodyOfWater",
						  "@id":"schema:BodyOfWater",
						  "layer":"core",
						  "children":[
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A canal, like the Panama Canal.",
								"name":"Canal",
								"@id":"schema:Canal",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A lake (for example, Lake Pontrachain).",
								"name":"LakeBodyOfWater",
								"@id":"schema:LakeBodyOfWater",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"An ocean (for example, the Pacific).",
								"name":"OceanBodyOfWater",
								"@id":"schema:OceanBodyOfWater",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A pond.",
								"name":"Pond",
								"@id":"schema:Pond",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A reservoir of water, typically an artificially created lake, like the Lake Kariba reservoir.",
								"name":"Reservoir",
								"@id":"schema:Reservoir",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A river (for example, the broad majestic Shannon).",
								"name":"RiverBodyOfWater",
								"@id":"schema:RiverBodyOfWater",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A sea (for example, the Caspian sea).",
								"name":"SeaBodyOfWater",
								"@id":"schema:SeaBodyOfWater",
								"layer":"core"
							 },
							 {
								"@type":"rdfs:Class",
								"rdfs:subClassOf":"schema:BodyOfWater",
								"description":"A waterfall, like Niagara.",
								"name":"Waterfall",
								"@id":"schema:Waterfall",
								"layer":"core"
							 }
						  ]
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Landform",
						  "description":"One of the continents (for example, Europe or Africa).",
						  "name":"Continent",
						  "@id":"schema:Continent",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Landform",
						  "description":"A mountain, like Mount Whitney or Mount Everest.",
						  "name":"Mountain",
						  "@id":"schema:Mountain",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Landform",
						  "description":"A volcano, like Fuji san.",
						  "name":"Volcano",
						  "@id":"schema:Volcano",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"An historical landmark or building.",
					"name":"LandmarksOrHistoricalBuildings",
					"@id":"schema:LandmarksOrHistoricalBuildings",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"The place where a person lives.",
					"name":"Residence",
					"@id":"schema:Residence",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Residence",
						  "description":"Residence type: Apartment complex.",
						  "name":"ApartmentComplex",
						  "@id":"schema:ApartmentComplex",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Residence",
						  "description":"Residence type: Gated community.",
						  "name":"GatedResidenceCommunity",
						  "@id":"schema:GatedResidenceCommunity",
						  "layer":"core"
					   }
					]
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"A tourist attraction. In principle any Thing can be a TouristAttraction, from a Mountain and LandmarksOrHistoricalBuildings to a LocalBusiness...",
					"name":"TouristAttraction",
					"@id":"schema:TouristAttraction",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Place",
					"description":"A tourist destination. In principle any Place can be a TouristDestination from a City, Region or Country to an AmusementPark or Hotel...",
					"name":"TouristDestination",
					"@id":"schema:TouristDestination",
					"layer":"pending"
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"Any offered product or service. For example: a pair of shoes; a concert ticket; the rental of a car; a haircut; or an episode of a TV show streamed online.",
			  "name":"Product",
			  "@id":"schema:Product",
			  "layer":"core",
			  "children":[
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Product",
					"description":"A single, identifiable product instance (e.g. a laptop with a particular serial number).",
					"name":"IndividualProduct",
					"@id":"schema:IndividualProduct",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Product",
					"description":"A datasheet or vendor specification of a product (in the sense of a prototypical description).",
					"name":"ProductModel",
					"@id":"schema:ProductModel",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Product",
					"description":"A placeholder for multiple similar products of the same kind.",
					"name":"SomeProducts",
					"@id":"schema:SomeProducts",
					"layer":"core"
				 },
				 {
					"@type":"rdfs:Class",
					"rdfs:subClassOf":"schema:Product",
					"description":"A vehicle is a device that is designed or used to transport people or cargo over land, water, air, or through space.",
					"name":"Vehicle",
					"@id":"schema:Vehicle",
					"layer":"core",
					"children":[
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Vehicle",
						  "description":"A bus (also omnibus or autobus) is a road vehicle designed to carry passengers...",
						  "name":"BusOrCoach",
						  "@id":"schema:BusOrCoach",
						  "layer":"auto"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Vehicle",
						  "description":"A car is a wheeled, self-powered motor vehicle used for transportation.",
						  "name":"Car",
						  "@id":"schema:Car",
						  "layer":"core"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Vehicle",
						  "description":"A motorcycle or motorbike is a single-track, two-wheeled motor vehicle.",
						  "name":"Motorcycle",
						  "@id":"schema:Motorcycle",
						  "layer":"auto"
					   },
					   {
						  "@type":"rdfs:Class",
						  "rdfs:subClassOf":"schema:Vehicle",
						  "description":"A motorized bicycle is a bicycle with an attached motor used to power the vehicle, or to assist with pedaling.",
						  "name":"MotorizedBicycle",
						  "@id":"schema:MotorizedBicycle",
						  "layer":"auto"
					   }
					]
				 }
			  ]
		   },
		   {
			  "@type":"rdfs:Class",
			  "rdfs:subClassOf":"schema:Thing",
			  "description":"A StupidType for testing.",
			  "name":"StupidType",
			  "@id":"schema:StupidType",
			  "layer":"attic"
		   }
		]
	 }
	`

	var obj interface{}
	var err = json.Unmarshal([]byte(dd), &obj)
	if err != nil {
		fmt.Println(err)
	}
}