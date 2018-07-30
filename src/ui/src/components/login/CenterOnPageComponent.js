import Grid from '@material-ui/core/Grid';
import React from 'react';

const CenterOnPage = WrappedComponent => props => (
  <Grid container>
    <Grid item xs={12}>
      <Grid container justify="center">
        <Grid item>
          <WrappedComponent {...props} />
        </Grid>
      </Grid>
    </Grid>
  </Grid>
);

export default CenterOnPage;
