import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import { SERVER, REQUEST } from "../../consts/consts";

import ServerConnectionForm from "../ServerConnectionForm/ServerConnectionForm";
import RequestForm from "../RequestForm/RequestForm";

import "./FormSwitcher.css";

export default function FormSwitcher() {
  const formSwitcherCn = new cn("FormSwitcher");
  const [activeTab, setActiveTab] = React.useState(SERVER);

  const onTabClick = (e) => {
    setActiveTab(e.target.textContent);

  };

  return (
    <div className={formSwitcherCn.root()}>
      <div className={formSwitcherCn.elem("tabs")}>
        {[SERVER, REQUEST].map((tab_title) => {
          return (
            <button
              className={formSwitcherCn
                .mix(formSwitcherCn.elem("tab", activeTab === tab_title ? "selected" : undefined))
                .elem("tab")}
              onClick={onTabClick}
            >
              {tab_title}
            </button>
          );
        })}
      </div>

      {activeTab === SERVER ?
        (<ServerConnectionForm/>) :
        (<RequestForm/>)
      }
    </div>
  );
}
