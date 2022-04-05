import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import "../CommonForm/CommonForm.css";
import "./RequestForm.css";

export default function RequestForm() {
  const requestFormCn = new cn("RequestForm");
  const commonFormCn = new cn("CommonForm");

  return (
    <form className={requestFormCn.mix(commonFormCn.root()).root()}>
      Select two Wikipedia pages, distance between which you wish to find
    </form>
  )
}
