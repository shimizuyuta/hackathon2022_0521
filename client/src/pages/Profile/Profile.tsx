import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import { FaDocker, FaVuejs, FaNodeJs, FaPython, } from 'react-icons/fa';
import { SiTypescript, SiRuby,  } from 'react-icons/si';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';

const Profile = () => {

  return (
  <Box sx={{p:2, m:3, mx: 'auto' }} style={{ maxWidth:'800px'}}>
  <Paper
    sx={{
      p: 2,
      m:3,
      margin: 'auto',
      maxWidth: 600,
      flexGrow: 1,
      mb:5,
    }}
    elevation={3}
    className='fafa'
  >
    <Grid
      container
      direction="row"
      justifyContent="center"
      alignItems="center"
    >
      <AccountCircleIcon sx={{ fontSize:130 }}/>
      <Stack spacing={1} sx={{ mx:3,px:3 }}>
        <h2 style={{ marginBottom:0 }}>清水優太</h2>
        <h5>example.com</h5>
      </Stack>
    </Grid>
  </Paper>
  <Paper
    sx={{
      p: 2,
      m:3,
      margin: 'auto',
      maxWidth: 600,
      flexGrow: 1,
    }}
    elevation={3}
  >
    <h2 style={{ paddingLeft:'2rem', display:'inline-block'}}>スキル</h2>
    <span 
      style={{
        marginLeft:5,
        paddingLeft:5,
        marginTop:3,
        paddingTop:5
      }}>
    </span>
    <Grid 
      container direction={'row'}  
      justifyContent="space-evenly" 
      alignItems="center"
    >
      <Grid direction={'column'}>
        <h3>TypeScript</h3>
        <SiTypescript size={100} />
      </Grid>
      <Grid direction={'column'}>
        <h3>Docker</h3>
        <FaDocker size={100} />
      </Grid>
      <Grid direction={'column'}>
        <h3>Vue.js</h3>
        <FaVuejs size={100} />
      </Grid>
      <Grid direction={'column'}>
        <h3>Node.js</h3>
        <FaNodeJs size={100} />
      </Grid>
    </Grid>
    <Grid container direction={'row'}
      justifyContent="space-evenly" 
      alignItems="center"
    >
      <Grid direction={'column'}>
        <h3>Ruby</h3>
        <SiRuby size={100} />
        </Grid>
      <Grid direction={'column'}>
        <h3>Python</h3>
        <FaPython size={100} />
      </Grid>
    </Grid>
  </Paper>
</Box>
  )
}

export default Profile