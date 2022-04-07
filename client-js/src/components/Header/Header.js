import * as React from "react";
import cn from "../../utils/CSSClassGenerator";

import "./Header.css";

export default function Header() {
    const headerCn = new cn("Header");

    return (
      <header className={headerCn.root()}>
        <a href="https://en.wikipedia.org/wiki/Main_Page" rel="noreferrer" target="_blank">
          <img
            src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b3/Wikipedia-logo-v2-en.svg/1200px-Wikipedia-logo-v2-en.svg.png"
            alt="Logo"
            className={headerCn.elem("logo")}
          />
        </a>
        <h1 className={headerCn.elem("title")}>
          Wiki Pages Distance Parser
        </h1>
      </header>
    )
}
