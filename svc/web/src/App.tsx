import { RouterProvider } from 'react-router-dom'
import router from './router'
import theme from './theme'
import { CssBaseline, ThemeProvider } from '@mui/material'

function App() {
  return (
    <>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <RouterProvider router={router} />
      </ThemeProvider>
    </>
  )
}

export default App
