import cn from "../../utils/CSSClassGenerator";
import Header from "../Header/Header";
import FormSwitcher from "../FormSwitcher/FormSwitcher";

import './App.css';

export default function App() {
  const appCn = new cn("App");

  return (
    <div className={appCn.root()}>
      <Header />
      <FormSwitcher />
    </div>
  );
}
