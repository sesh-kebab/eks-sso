import { connect } from 'react-redux';
import AddIAMUserModalComponent from './AddIAMUserModalComponent'
import { addCredentials } from '../../actions/aws-credential-actions'
import { hideIAMAddModal } from '../../actions/iam-modal-actions'

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