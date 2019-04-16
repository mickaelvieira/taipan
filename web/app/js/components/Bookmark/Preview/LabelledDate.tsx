import { wrap } from "dom-element-wrapper";

const LabelledDate = ({ label, date }) =>
  wrap("div").append(wrap("strong").append(label), wrap("span").append(date));

export default LabelledDate;
