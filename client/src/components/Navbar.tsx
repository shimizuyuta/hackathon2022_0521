import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
const Navbar = () => {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }} >
          <img src={`${process.env.PUBLIC_URL}/logo.png`} style={{ width:'10rem',backgroundColor:'white'}} alt="Logo" />
          </Typography>
          <Button color="inherit" href='/SignUp'>SignUp</Button>
          <Button color="inherit" href='/Login'>Login</Button>
        </Toolbar>
      </AppBar>
    </Box>
  )
}

export default Navbar