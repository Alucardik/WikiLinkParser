import * as WikiLinkParser from "../proto/WikiLinkParserService_pb";

function constructPublishRequest(initPage, targetPage) {
  const newPublishRequest = new WikiLinkParser.ParseRequest();
  newPublishRequest.setInitpage(initPage);
  newPublishRequest.setTargetpage(targetPage);

  return newPublishRequest;
}

function constructEmptyMsg() {
  return new WikiLinkParser.EmptyMsg();
}

export {
  constructPublishRequest,
  constructEmptyMsg
}
