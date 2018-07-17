import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { withTheme } from '@material-ui/core/styles';
import { authenticate } from '../../actions/authActions'
import { addCredentials } from '../../actions/awsCredentialActions'

import HomeComponent from './HomeComponent'

// todo: this component could be a pure functional component
const mapStateToProps = ({ user, auth, credentials }) => ({
  user,
  auth,
  credentials,
});

const mapActionsToProps = ({
  authenticate: authenticate,
  addCredentials: addCredentials,
});

export default withRouter(connect(
  mapStateToProps,
  mapActionsToProps
)(withTheme()(HomeComponent)));