import { connect } from 'react-redux';
import { authenticate } from '../../actions/authActions'
import { withStyles } from '@material-ui/core/styles';
import LoginComponent from './LoginComponent';


const styles = theme => ({
  root: theme.mixins.gutters({
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 16,
    width: 245,
  }),
  text: {
    width: '100%',
  },
  header: {
    backgroundColor: theme.palette.secondary.main,
    paddingTop: '56.25%', // 16:9,
    marginLeft: -theme.spacing.unit * 3,
    marginRight: -theme.spacing.unit * 3,
    padding: theme.spacing.unit * 3,
    marginBottom: theme.spacing.unit * 2,
    maxWidth: 245,
  }
});

const mapStateToProps = state => ({
  auth: state.auth,
  login: state.login,
});

const mapActionsToProps = ({
  authenticate: authenticate,
});

export default connect(
  mapStateToProps,
  mapActionsToProps,
)(withStyles(styles)(LoginComponent));