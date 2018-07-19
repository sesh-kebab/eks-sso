import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { addCredentials } from '../../actions/awsCredentialActions';
import { logoutUser } from '../../actions/authenticateActions';
import { showIAMAddModal } from '../../actions/iamAddUserModalActions';
import MainComponent from './MainComponent';

const mapStateToProps = state => ({
  auth: state.auth,
  credentials: state.credentials,
});

const mapActionsToProps = {
  logout: logoutUser,
  addCredentials: addCredentials,
  showIAMAddModal: showIAMAddModal,
};

export default withRouter(
  connect(
    mapStateToProps,
    mapActionsToProps
  )(MainComponent)
);
