import React, { useState } from 'react'
import { Button, Container, Stack, TextField, Box } from '@mui/material'
import { SubmitHandler, useForm } from 'react-hook-form'

type LoginUserForm = {
  email:string,
  password:string,
  name:string,
}

const Login = () => {
  const { register, handleSubmit } = useForm<LoginUserForm>()
  const onSubmit:SubmitHandler<LoginUserForm> = (data) =>{
    console.log(data)
  }

  return (
    <Container maxWidth="sm" sx={{ pt: 5 ,mt:5}}>
      <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>
        <Stack spacing={3}>
          <TextField 
            required
            label="メールアドレス"
            type="email"
            {...register('email')}
          />
          <TextField
            required
            label="パスワード" 
            type="password" 
            {...register('password')}
          />
          <Button 
            color="primary"
            variant="contained" 
            size="large"
            type='submit'
          >
            作成
          </Button>
        </Stack>
      </Box>
    </Container>
  )
}

export default Login
