import CssBaseline from "@material-ui/core/CssBaseline";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";

import useStyle from "./style";

function Tasks() {
	const classes = useStyle();
	return (
		<div className={classes.root}>
			<CssBaseline />
			<Container maxWidth="lg" className={classes.container}>
				<Grid container spacing={3}>
					<Grid item xs={12}>
						<Paper className={classes.paper}>
							<h1>Hello</h1>
						</Paper>
					</Grid>
				</Grid>
			</Container>
		</div>
	);
}
export default Tasks;
