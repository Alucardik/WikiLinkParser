import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import { WikiLinkParserClient } from "../../proto/WikiLinkParserService_grpc_web_pb";
import {PROXY_HOSTNAME} from "../../config";

import Header from "../Header/Header";
import FormSwitcher from "../FormSwitcher/FormSwitcher";


import './App.css';
import {constructEmptyMsg, constructPublishRequest} from "../../utils/GRPCutils";

export default function App() {
  const appCn = new cn("App");
  const grpcClient = new WikiLinkParserClient(PROXY_HOSTNAME, null, null);

  const PublishTask = (initPage, targetPage) => {
    grpcClient.publishTask(constructPublishRequest(initPage, targetPage), null, (err, resp) => {
      if (err) {
        return console.log(err);
      }

      const error = resp.getError();
      const msg = resp.getMsg();

      if (error === 1) {
        console.log(error, msg);
      }

      console.log(msg);
    });
  };

  const Connect = () => {
    grpcClient.establishConnection(constructEmptyMsg(), null, (err, resp) => {
      if (err) {
        return console.log(err);
      }

      const error = resp.getError();
      const msg = resp.getMsg();

      console.log(error, msg);
    });
  };

  return (
    <div className={appCn.root()}>
      <Header />
      <FormSwitcher onPublishSubmit={PublishTask} onConnectSubmit={Connect} />
    </div>
  );
}
