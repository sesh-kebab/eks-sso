import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import { logout } from '../../actions/auth-actions'
import { addCredentials } from '../../actions/aws-credential-actions'
import { showIAMAddModal } from '../../actions/iam-modal-actions'
import MainComponent from "./MainComponent";

const drawerWidth = 240;
const styles = theme => ({
  root: {
      flexGrow: 1,
      zIndex: 1,
      overflow: 'hidden',
      position: 'relative',
      display: 'flex',
  },
  flex: {
      flex: 1,
  },
  appBar: {
      zIndex: theme.zIndex.drawer + 1,
  },
  drawerPaper: {
      position: 'relative',
      width: drawerWidth,
  },
  content: {
      flexGrow: 1,
      backgroundColor: theme.palette.background.default,
      padding: theme.spacing.unit * 3,
      minWidth: 0, // So the Typography noWrap works
  },
  toolbar: theme.mixins.toolbar,
});

const mapStateToProps = state => ({
  auth: state.auth,
  credentials: state.credentials,
});

const mapActionsToProps = ({
  logout: logout,
  addCredentials: addCredentials,
  showIAMAddModal: showIAMAddModal,
});

export default withRouter(connect(
  mapStateToProps,
  mapActionsToProps
)(withStyles(styles, { withTheme: true })(MainComponent)));