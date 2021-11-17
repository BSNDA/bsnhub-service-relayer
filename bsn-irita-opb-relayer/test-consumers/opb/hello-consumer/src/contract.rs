use cosmwasm_std::{attr,entry_point, Binary, Deps, DepsMut, Env, Addr, Response, MessageInfo, StdResult,WasmMsg, SubMsg, to_binary};
use json::object;
use base64::encode;

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg,QueryMsg};
use crate::state::{State,config};

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InitMsg,
) -> Result<Response, ContractError>  {
    let state = State {
        owner: deps.api.addr_canonicalize(&info.sender.to_string())?,
        core_contract: msg.core_contract,
        target_contract: msg.target_contract,
    };
    config(deps.storage).save(&state)?;

    Ok(Response::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    env: Env,
    _info: MessageInfo,
    msg: HandleMsg,
) -> Result<Response, ContractError>  {
    match msg {
        HandleMsg::Hello{words} => try_hello(deps,words,env.contract.address),
        HandleMsg::CallBack{request_id, words} => call_back(deps, request_id, words),
    }
}

pub fn try_hello(deps: DepsMut, words: String, self_address: Addr) -> Result<Response, ContractError>  {
    let state = config(deps.storage).load()?;
    let endpoint_address = state.target_contract.as_str();
    let self_address_str = self_address.as_str();
    let endpoint_info = object!{"dest_chain_id"=>"1","dest_chain_type"=>"opb","endpoint_address"=>endpoint_address,"endpoint_type"=>"contract"};
    let call_data = object!{"hello"=>object! {"words"=>words}};
    let call_data_base64 = encode(call_data.to_string().as_bytes());
    let msg = object!{"send_request"=> object!{"endpoint_info"=>endpoint_info.to_string(),"method"=>"hello","call_data"=>call_data_base64,"callback_address"=>self_address_str,"callback_function"=>"call_back"}};

    let messages =vec![SubMsg::new(WasmMsg::Execute {
        contract_addr: state.core_contract.to_string(),
        msg:  Binary::from(msg.to_string().as_bytes()),
        funds: vec![],
    })];

    let mut res = Response::default();
    res.messages = messages;

    Ok(res)
}

pub fn call_back(_deps: DepsMut, request_id: String, words: String) -> Result<Response, ContractError> {
    let mut res = Response::default();
    res.attributes = vec![attr("request_id",request_id),attr("words",words)];
    Ok(res)
}

pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    to_binary("no query function")
}
