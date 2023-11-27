import { encoder } from "./keys.js";
import bls from "@chainsafe/bls";
import { keys } from "../constants.js";

export const sign = (data) => {
  const json = JSON.stringify(data);
  const buffer = Buffer.from(json, "utf8");
  return encoder.encode(keys.secretKey.sign(buffer).toBytes());
};

export const attest = (data, signature, signers) => {
  const signer = encoder.encode(keys.publicKey.toBytes());
  if (signers.includes(signer)) {
    return null;
  }
  const soloSignature = sign(data);
  const aggregatedSignature = signature
    ? encoder.encode(
        bls.Signature.aggregate([
          bls.Signature.fromBytes(Buffer.from(encoder.decode(soloSignature))),
          bls.Signature.fromBytes(Buffer.from(encoder.decode(signature))),
        ]).toBytes()
      )
    : soloSignature;
  return {
    data,
    signature: aggregatedSignature,
    signers: [...signers, signer],
  };
};