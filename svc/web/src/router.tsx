import { createBrowserRouter, redirect } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import SignupPage from "@pages/SignupPage";
import CreateAccountPage from "@pages/CreateAccountPage";

const router = createBrowserRouter([
  {
    path: "/jared-hughes",
    element: <LandingPage />
  },
  {
    path: "/signup",
    element: <SignupPage />
  },
  {
    path: "/accounts/activate",
    element: <CreateAccountPage />
  },
  {
    path: "*",
    loader: () => redirect("/jared-hughes"),
  }
]);

export default router;