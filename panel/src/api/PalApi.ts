import request from "../util/request"
import {
  RebootParam,
  BroadcastParam,
  PalPlayer,
  MemStatus,
  RconParam,
  SimpleParam,
  ServerStateResponse,
} from "../model/PalApiModel"

export function getMemoryStatus(): Promise<MemStatus> {
  return request.get(`/api/memory`)
}

export function getOnlinePlayers(): Promise<Array<PalPlayer>> {
  return request.get(`/api/players`)
}

export function broadcast(param: BroadcastParam): Promise<void> {
  return request.post(`/api/broadcast`, param)
}

export function rcon(param: RconParam): Promise<SimpleParam> {
  return request.post(`/api/rcon`, param)
}

export function reboot(param: RebootParam): Promise<void> {
  return request.post(`/api/reboot`, param)
}

export function getServerState(): Promise<ServerStateResponse> {
  return request.get(`/api/state`)
}