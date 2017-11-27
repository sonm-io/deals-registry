# Deals 


## Phase I  - Version 1.0 - version MVP
Deals are time-limited. It means, that a deal has an exact time term.

## Struct ``Deal``

| Name  | Type  | Description  |
|---|---|---|
| orderHash  | uint | SHA-256 of the ASK order hash.  |
| client  | address  | Ethereum addr of the consumer.  |
| hub | address  | Ethereum addr of the consumer (hub owner).  |
| price | uint256  | Price in SNM tokens, in "wei" units. |
| startTime | uint  | Start time in Unix time (EPOCH). |
| workTime | uint |  Deal time in seconds. |
| endTime | uint |  End time in Unix time (EPOCH). |
| status | uint  |  see ``Status disclosure`` |

## ```Status``` disclosure
| Name  | Value  | Description  |
|---|---|---|
| any  | 0 | is unreachable in blockchain implementation  |
| Opened  | 1 | created by buyer and waiting accept from supplier  |
| Accepted  | 2 | accepted by supplier, ready to work  |
| Closed  | 3 | closed by any reasons  |


**Dispute**

No dispute mechanics.


**Part time split**

Yes. On the consumer's will.

***Rating***

No rating system.

***Network Fee***

No network fee.