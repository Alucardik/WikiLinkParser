import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import { SERVER, REQUEST } from "../../consts/consts";

import ServerConnectionForm from "../ServerConnectionForm/ServerConnectionForm";
import RequestForm from "../RequestForm/RequestForm";

import "./FormSwitcher.css";

export default function FormSwitcher(props) {
  const formSwitcherCn = new cn("FormSwitcher");
  const [activeTab, setActiveTab] = React.useState(SERVER);

  const onTabClick = (e) => {
    setActiveTab(e.target.textContent);
  };

  return (
    <div className={formSwitcherCn.root()}>
      <div className={formSwitcherCn.elem("tabs")}>
        {[SERVER, REQUEST].map((tab_title, index) => {
          return (
            <button key={index}
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
        (<ServerConnectionForm onSubmit={props.onConnectSubmit} />) :
        (<RequestForm onSubmit={props.onPublishSubmit} />)
      }
    </div>
  );
}
