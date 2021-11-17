
use std::u64;
use json::object;
use sha2::{Sha256, Digest};

use cosmwasm_std::{
    attr,to_binary,entry_point,Addr, Binary, Deps, DepsMut, Env, Response, MessageInfo, StdResult, WasmMsg, SubMsg
};

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg, QueryMsg};
use crate::state::{CHAINID, RELAYER, REQUESTS, REQUESTSEQUENCE, CALLBACKS,Callback};

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InitMsg,
) -> Result<Response, ContractError> {
    let mut relayer = msg.admin;

    if relayer.is_none(){
        relayer = Some(info.sender.clone());
    }

    let sequence: u64 = 0;
    REQUESTSEQUENCE.save(deps.storage, &sequence)?;
    CHAINID.save(deps.storage,&msg.source_chain_id)?;
    RELAYER.set(deps,relayer)?;

    Ok(Response::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> Result<Response, ContractError>  {
    match msg {

        HandleMsg::SendRequest {endpoint_info, method, call_data, callback_address, callback_function}
        => send_request(deps,endpoint_info, method, call_data, callback_address, callback_function),
        HandleMsg::SetResponse { request_id,err_msg, output } => set_response(deps, request_id,err_msg, output),
        HandleMsg::SetRelayer{ relayer } => set_relayer(deps, relayer,info.sender),
    }
}

pub fn send_request(deps: DepsMut,endpoint_info: String, _method: String, call_data: Binary, callback_address: Addr, call_back_function: String) -> Result<Response, ContractError> {
    let chain_id = CHAINID.load(deps.storage)?;
    let mut sequence = REQUESTSEQUENCE.load(deps.storage)?;

    let request_id_str = chain_id+&sequence.to_string();

    let mut hasher = Sha256::new();
    hasher.input(&request_id_str.into_bytes());
    let output = hasher.result();
    let request_id = format!("{:X}",output);

    sequence += 1;
    REQUESTSEQUENCE.save(deps.storage, &sequence)?;

    let call_back = Callback{
        address: callback_address,
        method: call_back_function,
    };
    CALLBACKS.save(deps.storage, &request_id, &call_back)?;
    let mut res = Response::default();
    res.attributes = vec![attr("request_id", request_id),attr("endpoint_info", endpoint_info),attr("method", _method),attr("callData", call_data.to_string())];
    Ok(res)
}

pub fn set_response(deps: DepsMut,request_id: String, err_msg: String, output: String) -> Result<Response, ContractError>  {
    let executed = REQUESTS.load(deps.storage, &request_id);
    if executed.is_ok() {
        Err(ContractError::Unauthorized{})
    } else {
        let mut result = err_msg;
        if output.len()>0 {
            result = output;
        }

        REQUESTS.save(deps.storage, &request_id.clone(), &true)?;

        let callback = CALLBACKS.load(deps.storage, &request_id.clone())?;
        let result_obg = object!{ callback.method.as_str() => object!{"request_id"=>request_id.clone(),"words"=>result}};

        let messages =vec![SubMsg::new(WasmMsg::Execute {
            contract_addr: callback.address.to_string(),
            msg: Binary::from(result_obg.to_string().as_bytes()),
            funds: vec![],
        })];

        REQUESTS.save(deps.storage, &request_id.clone(), &true)?;
        let mut res = Response::default();
        res.messages = messages;
        Ok(res)
    }
}

pub fn set_relayer(deps: DepsMut, relayer: Option<Addr>, caller:Addr) ->  Result<Response, ContractError>  {
    RELAYER.assert_admin(deps.as_ref(), &caller);
    RELAYER.set(deps, relayer);
    Ok(Response::default())
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::GetReuqest {requst_id} => to_binary(&query_request(deps,requst_id)?),
        QueryMsg::GetSequence {} => to_binary(&query_sequence(deps)?),
        QueryMsg::GetRelayer {} => to_binary( &RELAYER.query_admin(deps)?),
    }
}

fn query_request(deps: Deps,requst_id: String) -> StdResult<bool> {
    let state = REQUESTS.load(deps.storage,&requst_id)?;
    Ok(state)
}

fn query_sequence(deps: Deps) -> StdResult<u64> {
    let state = REQUESTSEQUENCE.load(deps.storage)?;
    Ok(state)
}
