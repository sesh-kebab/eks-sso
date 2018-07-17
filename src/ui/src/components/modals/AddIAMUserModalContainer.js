import { connect } from 'react-redux';
import AddIAMUserModalComponent from './AddIAMUserModalComponent'
import { addCredentials } from '../../actions/awsCredentialActions'
import { hideIAMAddModal } from '../../actions/iamAddUserModalActions'

const mapStateToProps = state => ({
  credentials: state.credentials,
  iamModal: state.iamModal,
});

const mapActionsToProps = ({
  addCredentials: addCredentials,
  hideModal: hideIAMAddModal,
});

export default connect(
  mapStateToProps,
  mapActionsToProps
)(AddIAMUserModalComponent);