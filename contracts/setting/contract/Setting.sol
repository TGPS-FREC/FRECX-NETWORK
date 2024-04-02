pragma solidity ^0.4.21;

import "./libs/SafeMath.sol";

contract Setting {
    using SafeMath for uint256;

    address public owner;
    uint64 public epochValue;
    uint64 public maxMasterNode;
    uint64 public rewardPerEpoch;
    uint64 public TIPIncreaseMasternodes;

    modifier onlyValidCandidate {
        require(msg.sender == owner);
        _;
    }

    event epoch(address _modifiedby,uint64 _value);
    event reward(address _modifiedby,uint64 _value);
    event masternode(address _modifiedby,uint64 _value);
    event tipIncreaseMasterNode(address _modifiedby,uint64 _value);

    constructor(
        uint64 _maxMasterNode,
        uint64 _epochValue,
        uint64 _rewardPerEpoch,
        uint64 _TIPIncreaseMasternodes
    ) public {
        owner = msg.sender;
        epochValue = _epochValue;
        maxMasterNode = _maxMasterNode;
        rewardPerEpoch = _rewardPerEpoch;
        TIPIncreaseMasternodes = _TIPIncreaseMasternodes;
    }

    function getEpoch() public view returns(uint64){
        return epochValue;
    }

    function getRewardPerEpoch() public view returns(uint64){
        return rewardPerEpoch;
    }

    function getMaxMasterNode() public view returns(uint64){
        return maxMasterNode;
    }

    function getTIPIncrease() public view returns(uint64){
        return TIPIncreaseMasterNodes;
    }


    function setEpoch(uint64 _val) onlyValidCandidate public{
        epochValue = _val;
        emit epoch(msg.sender, _val);
    }

    function setRewardPerEpoch(uint64 _val) onlyValidCandidate public{
        rewardPerEpoch = _val;
        emit reward(msg.sender, _val);

    }

    function setMaxMasterNode(uint64 _val)onlyValidCandidate public{
        maxMasterNode = _val;
        emit masternode(msg.sender, _val);

    }

    function setTIPIncrease(uint64 _val)onlyValidCandidate public{
        TIPIncreaseMasterNodes = _val;
        emit tipIncreaseMasterNode(msg.sender, _val);

    }

}
