import { compose, setDisplayName, withStateHandlers } from 'recompose';
import { connect } from 'react-redux';

import { addCredentials } from '../../actions/awsCredentialActions';
import { hideIAMAddModal } from '../../actions/iamAddUserModalActions';
import AddIAMUserModalComponent from './AddIAMUserModalComponent';

const mapStateToProps = state => ({
  iamModal: state.iamModal,
});

const mapActionsToProps = {
  addCredentials: addCredentials,
  hideModal: hideIAMAddModal,
};

const enhance = compose(
  setDisplayName('AddIamUserModal'),
  connect(
    mapStateToProps,
    mapActionsToProps
  ),
  withStateHandlers(
    () => ({
      accessId: '',
      secretKey: '',
    }),
    {
      onAccessIdChange: state => event => {
        state.accessId = event.target.value;
        return state;
      },
      onSecretKeyChange: state => event => {
        state.secretKey = event.target.value;
        return state;
      },
      onClose: (state, props) => () => {
        props.hideModal();
      },
      onSubmit: (state, props) => event => {
        event.preventDefault();
        props.addCredentials(state.accessId, state.secretKey);
      },
    }
  )
);

export default enhance(AddIAMUserModalComponent);
