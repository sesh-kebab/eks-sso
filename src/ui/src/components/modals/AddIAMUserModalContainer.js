import { connect } from 'react-redux';
import { addCredentials } from '../../actions/awsCredentialActions'
import { hideIAMAddModal } from '../../actions/iamAddUserModalActions';
import AddIAMUserModalComponent from './AddIAMUserModalComponent';

const mapStateToProps = state => ({
  credentials: state.credentials,
  iamModal: state.iamModal,
});

const mapActionsToProps = {
  addCredentials: addCredentials,
  hideModal: hideIAMAddModal,
};

export default connect(
  mapStateToProps,
  mapActionsToProps
)(AddIAMUserModalComponent);
