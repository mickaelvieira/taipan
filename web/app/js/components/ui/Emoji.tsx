import React from "react";
import { Emoji as EmojiBase } from "emoji-mart/dist-modern/index";

interface Props {
  emoji: string;
  size?: number;
}

export default function Emoji({ emoji, ...rest }: Props): JSX.Element {
  return <EmojiBase emoji={emoji} sheetSize={32} size={32} {...rest} />;
}
