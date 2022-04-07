import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import "../CommonForm/CommonForm.css";
import "./RequestForm.css";

export default function RequestForm(props) {
  const requestFormCn = new cn("RequestForm");
  const commonFormCn = new cn("CommonForm");

  const onSubmit = (e) => {
    e.preventDefault();
    const initPage = "https://en.wikipedia.org/wiki/Main_Page";
    const targetPage = "https://en.wikipedia.org/wiki/Viatkogorgon";

    props.onSubmit(initPage, targetPage);
  };

  return (
    <form
      className={requestFormCn.mix(commonFormCn.root()).root()}
      name="server_credentials"
      onSubmit={onSubmit}
    >
      <h2 className={commonFormCn.elem("title")}>
        Enter server credentials to connect
      </h2>

      <label className={commonFormCn.elem("label")}>
        Start page:
        <input
          name="init_page"
          type="text"
          placeholder="https://en.wikipedia.org/wiki/Main_Page"
          className={commonFormCn.elem("input")}
        />
      </label>

      <label className={commonFormCn.elem("label")}>
        Target page:
        <input
          name="target_page"
          type="text"
          placeholder="https://en.wikipedia.org/wiki/Main_Page"
          className={commonFormCn.elem("input")}
        />
      </label>

      <button
        type="submit"
        className={commonFormCn.elem("submit")}
      >
        Send request
      </button>
    </form>
  );
}
