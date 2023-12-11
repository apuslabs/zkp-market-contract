
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// import "./console.sol";
import "./ApusData.sol";
import "./market.sol";

contract QueryProxy {

    Market private market;
    Task private

    constructor(address marketAddress) {
        market = Market(marketAddress);
        // console.log("market init");
    }

    function getClientCount() public view returns(uint256){
        uint256 total = 0;
        for(uint i = 0; i < market.getClientCount(); i++) {
            try market.clients(i) {
                (
                    , // address owner
                    , // uint256 id
                    , // string memory url
                    , // uint256 minFee
                    uint8 maxZkEvmInstance,
                    , // uint8 curInstance
                    ApusData.ClientStatus stat
                ) = market.clients(i);
                if (stat == ApusData.ClientStatus.Running) {
                    total += maxZkEvmInstance;
                }
            } catch {
                break;
            }
        }
        return total;
    }

    function getAvilableClientCount() public view returns(uint256){
        uint256 total = 0;
        for(uint i = 0; i < market.getClientCount(); i++) {
            try market.clients(i) {
                (
                    , // address owner,
                    , // uint256 id,
                    , // string memory url,
                    , // uint256 minFee,
                    uint8 maxZkEvmInstance,
                    uint8 curInstance,
                    ApusData.ClientStatus stat
                ) = market.clients(i);
                if (stat == ApusData.ClientStatus.Running && curInstance < maxZkEvmInstance) {
                    total += maxZkEvmInstance - curInstance;
                }
            } catch {
                break;
            }
        }
        return total;
    }

    function getAllClient() public view returns(ApusData.ClientConfig[] memory){
        ApusData.ClientConfig[] memory clients = new ApusData.ClientConfig[](market.getClientCount());
        for(uint i = 0; i < market.getClientCount(); i++) {
            try market.clients(i) {
                (
                    address owner,
                    uint256 id,
                    string memory url,
                    uint256 minFee,
                    uint8 maxZkEvmInstance,
                    uint8 curInstance,
                    ApusData.ClientStatus stat
                ) = market.clients(i);
                clients[i] = ApusData.ClientConfig(owner, id, url, minFee, maxZkEvmInstance, curInstance, stat);
            } catch {
                break;
            }
        }
        return clients;
    }

}