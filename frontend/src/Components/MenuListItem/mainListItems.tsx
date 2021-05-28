import React from "react";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import AssignmentTurnedInIcon from "@material-ui/icons/AssignmentTurnedIn";

export const mainListItems = (
	<div>
		<ListItem button>
			<ListItemIcon>
				<AssignmentTurnedInIcon />
			</ListItemIcon>
			<ListItemText primary="Tasks" />
		</ListItem>
	</div>
);
