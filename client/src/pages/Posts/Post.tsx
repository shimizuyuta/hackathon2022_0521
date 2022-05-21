import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper'
import CreateIcon from '@mui/icons-material/Create';
import React, { useState } from 'react';
import dayjs from 'dayjs';
import { ChangeEvent } from 'react';

type PostForm = {
  title: string,
  description: string,
  startDate:Date,
  endDate:Date,
}
const initalData: PostForm={
  title: '',
  description: '',
  startDate:new Date(),
  endDate:new Date(),
}

const Post = () => {
  const [post,setPost] = useState<PostForm>(initalData);
  const today = dayjs().format('YYYY-MM-DD')


  const onChangeTitle = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) =>{
    console.log(event.target.value)
    const title = event.target.value;
    setPost({...post,title:title})
  };
  const onChangeDescription = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) =>{
    const description = event.target.value;
    setPost({...post,description:description})
  }

  const onChangeStartDate = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) =>{
    const startDate: Date = new Date(event.target.value);
    setPost({...post,startDate:startDate })
  }
  const onChangeEndDate = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) =>{
    const endDate: Date = new Date(event.target.value);
    setPost({...post,endDate:endDate })
  }
  console.log(post,' post data')

  return (
    <Paper
      sx={{
        p: 2,
        m: 5,
        maxWidth: '900px',
        mx:'auto'
      }}
      elevation={3} 
    >
      <Grid container direction="column" >
        <Box sx={{ mx:3, py:1 }}>
          <Typography variant="h5" sx={{ py:1 }}>
            タイトル
          </Typography>
          <TextField 
            id="outlined-basic" 
            fullWidth 
            variant="outlined" 
            placeholder="タイトルを記述してください"
            value={post.title}
            onChange={e => onChangeTitle(e)}
          />
        </Box>
        <Box sx={{ mx:3, py:1 }}>
          <Typography variant="h5" sx={{ py:1 }}>
            募集期間
          </Typography>
          <TextField 
            variant="outlined" 
            label="開始日時"
            type="date"
            defaultValue={today}
            margin="normal"
            required
            style ={{width: '40%', marginRight: '20px'}} 
            // value={post.startDate}
            onChange={e => onChangeStartDate(e)}
          />
          <TextField 
            variant="outlined" 
            label="終了日時"
            type="date"
            defaultValue={today}
            margin="normal"
            required
            style ={{width: '40%', marginLeft: '20px'}} 
            // value={post.endDate}
            onChange={e => onChangeEndDate(e)}
          />
        </Box>
        <Box sx={{ mx:3, py:1, pb:4 }}>
          <Typography variant="h5" sx={{ py:1 }}>説明</Typography>
          <TextField 
            fullWidth
            id="fullWidth"
            placeholder="説明を記述してください"
            minRows='4'
            maxRows='8'
            multiline  
            value={post.description}
            onChange={e => onChangeDescription(e)}
          />
        </Box>
        <Button 
          variant="contained" 
          endIcon={<CreateIcon />}  
          size="large" 
          sx={{ mx:3, px:3 }}
        >
          投稿する
        </Button>
      </Grid>
    </Paper>

  )
}

export default Post