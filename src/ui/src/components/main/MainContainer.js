import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { hideIAMAddModal, showIAMAddModal } from '../../actions/iamAddUserModalActions';
import { logoutUser } from '../../actions/authenticateActions';
import MainComponent from './MainComponent';

const mapStateToProps = ({ auth, credentials }) => ({
  auth,
  credentials,
});

const mapDispatchToProps = {
  logoutUser,
  showIAMAddModal,
  hideIAMAddModal,
};

export default withRouter(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(MainComponent)
);
