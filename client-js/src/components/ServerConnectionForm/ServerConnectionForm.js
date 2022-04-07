import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import "../CommonForm/CommonForm.css";
import "./ServerConnectionForm.css";

export default function ServerConnectionForm(props) {
  const serverConnectionFormCn = new cn("ServerConnectionForm");
  const commonFormCn = new cn("CommonForm");

  // TODO: add state vars for inputs and set them to docker config by default

  const onSubmit = (e) => {
    e.preventDefault();

    props.onSubmit();
  }

  return (
    <form
      className={serverConnectionFormCn.mix(commonFormCn.root()).root()}
      name="server_credentials"
      onSubmit={onSubmit}
    >
      <h2 className={commonFormCn.elem("title")}>
        Enter server credentials to connect
      </h2>

      <label className={commonFormCn.elem("label")}>
        Server address:
        <input
          name="server_address"
          type="text"
          placeholder="localhost"
          className={commonFormCn.elem("input")}
        />
      </label>

      <label className={commonFormCn.elem("label")}>
        Server port:
        <input
          name="server_port"
          type="number"
          placeholder="3000"
          min={1}
          max={65535}
          className={commonFormCn.elem("input")}
        />
      </label>

      <button
        type="submit"
        className={commonFormCn.elem("submit")}
      >
        Connect
      </button>
    </form>
  );
}
