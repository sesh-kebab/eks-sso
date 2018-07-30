import { compose, setDisplayName, withStateHandlers } from 'recompose';
import { connect } from 'react-redux';

import { addCredentials } from '../../actions/awsCredentialActions';
import { hideIAMAddModal } from '../../actions/iamAddUserModalActions';
import AddIAMUserModalComponent from './AddIAMUserModalComponent';

const mapStateToProps = ({ iamModal }) => ({
  iamModal,
});

const mapDispatchToProps = {
  addCredentials,
  hideIAMAddModal,
};

const enhance = compose(
  setDisplayName('AddIamUserModal'),
  connect(
    mapStateToProps,
    mapDispatchToProps
  ),
  withStateHandlers(
    () => ({
      accessId: '',
      secretKey: '',
    }),
    {
      onAccessIdChange: state => event => ({ ...state, accessId: event.target.value }),
      onSecretKeyChange: state => event => ({ ...state, secretKey: event.target.value }),
      onSubmit: ({ accessId, secretKey }, props) => event => {
        event.preventDefault();
        props.addCredentials(accessId, secretKey);
      },
    }
  )
);

export default enhance(AddIAMUserModalComponent);
