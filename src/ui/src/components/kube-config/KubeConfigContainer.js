import { compose, lifecycle } from 'recompose';
import KubeConfigComponent from './KubeConfigComponent';
import api from './../../api'

const enhance = compose(
  // todo: add logic to show a loading indicator?
  lifecycle({
    state: { loading: true },
    componentDidMount() {
      api.getData('/cluster')
        .then(o => {
          return o.json();
        })
        .then(j => {
          this.setState({
            loading: false,
            kubeconfig: j.kubeconfig
          });
        })
        .catch(e => {
          console.error(e);
        });
    }
  })
);

export default enhance(KubeConfigComponent)