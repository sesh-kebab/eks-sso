import Grid from '@material-ui/core/Grid';
import React from 'react';

const centerOnPage = WrappedComponent => {
  class CenterOnPageComponent extends React.Component {
    render() {
      return (
        <Grid container>
          <Grid item xs={12}>
            <Grid container justify="center">
              <Grid item>
                <WrappedComponent {...this.props} />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      );
    }
  }

  CenterOnPageComponent.displayName = 'CenterOnPageComponent';

  return CenterOnPageComponent;
};

export default centerOnPage;
