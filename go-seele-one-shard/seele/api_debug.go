/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package seele

 import (
 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"runtime"
// 	"runtime/pprof"

// 	"github.com/seeleteam/go-seele/common"
 	"github.com/seeleteam/go-seele/core/types"
)

 // PrivateDebugAPI provides an API to access full node-related information for debug.
 type PrivateDebugAPI struct {
 	s *SeeleService
 }

 // NewPrivateDebugAPI creates a new NewPrivateDebugAPI object for rpc service.
 func NewPrivateDebugAPI(s *SeeleService) *PrivateDebugAPI {
 	return &PrivateDebugAPI{s}
 }

 // PrintBlock retrieves a block and returns its pretty printed form, when height is -1 the chain head is returned
 func (api *PrivateDebugAPI) PrintBlock(height int64) (*types.Block, error) {
 	block, err := getBlock(api.s.chains[0], height)
 	if err != nil {
 		return nil, err
 	}

 	return block, nil
}

// // GetTxPoolContent returns the transactions contained within the transaction pool
// func (api *PrivateDebugAPI) GetTxPoolContent() (map[string][]map[string]interface{}, error) {
// 	txPool := api.s.TxPool()
// 	data := txPool.GetTransactions(false, true)

// 	content := make(map[string][]map[string]interface{})
// 	for _, tx := range data {
// 		key := tx.Data.From.ToHex()
// 		content[key] = append(content[key], PrintableOutputTx(tx))
// 	}

// 	return content, nil
// }

// // GetTxPoolTxCount returns the number of transaction in the pool
// func (api *PrivateDebugAPI) GetTxPoolTxCount() (uint64, error) {
// 	txPool := api.s.TxPool()
// 	return uint64(txPool.GetPendingTxCount()), nil
// }

// // GetPendingTransactions returns all pending transactions
// func (api *PrivateDebugAPI) GetPendingTransactions() ([]map[string]interface{}, error) {
// 	pendingTxs := api.s.TxPool().GetTransactions(true, true)
// 	transactions := make([]map[string]interface{}, 0)
// 	for _, tx := range pendingTxs {
// 		transactions = append(transactions, PrintableOutputTx(tx))
// 	}

// 	return transactions, nil
// }

// // GetPendingDebts returns all pending debts
// // func (api *PrivateDebugAPI) GetPendingDebts() ([]*types.Debt, error) {
// // 	return api.s.DebtPool().GetAll(), nil
// // }

 // TpsInfo tps detail info
 type TpsInfo struct {
 	StartHeight uint64
 	EndHeight   uint64
 	Count       uint64
 	Duration    uint64
 }

 // GetTPS get tps info
 func (api *PrivateDebugAPI) GetTPS() (*TpsInfo, error) {
 	chains := api.s.BlockChain()
 	block := chains[0].CurrentBlock()
 	timeInterval := uint64(150)
 	if block.Header.Height == 0 {
 		return nil, nil
 	}

 	var count = uint64(len(block.Transactions) - 1)
 	var duration uint64
 	var endHeight uint64
 	startTime := block.Header.CreateTimestamp.Uint64()
 	for height := block.Header.Height - 1; height > 0; height-- {
 		current, err := chains[0].GetStore().GetBlockByHeight(height)
 		if err != nil {
 			return nil, fmt.Errorf("failed to get block, error:%s, block height:%d", err, height)
 		}

 		count += uint64(len(current.Transactions) - 1)
 		duration = startTime - current.Header.CreateTimestamp.Uint64()
 		endHeight = height

 		if duration > timeInterval {
 			break
 		}
 	}

 	return &TpsInfo{
 		StartHeight: endHeight,
 		EndHeight:   block.Header.Height,
 		Count:       count,
 		Duration:    duration,
 	}, nil
 }

// // DumpHeap dumps the heap usage.
// func (api *PrivateDebugAPI) DumpHeap(fileName string, gcBeforeDump bool) (string, error) {
// 	if len(fileName) == 0 {
// 		fileName = "heap.dump"
// 	}

// 	if gcBeforeDump {
// 		runtime.GC()
// 	}

// 	flie := filepath.Join(common.GetDefaultDataFolder(), fileName)
// 	f, err := os.Create(flie)
// 	if err != nil {
// 		return "", err
// 	}

// 	return flie, pprof.WriteHeapProfile(f)
// }
