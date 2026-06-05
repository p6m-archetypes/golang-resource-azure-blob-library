-- golang-resource-azure-blob-library main module.
-- Renders Azure Blob Storage client setup into the service's storage package.
--
-- The calling archetype is responsible for adding the corresponding
-- Go module dependency:
--   github.com/Azure/azure-sdk-for-go/sdk/storage/azblob
--
-- API (called from a parent archetype):
--   local azure = require("golang-resource-azure-blob")
--   azure.render(context, { destination = context:get("project-name") })

local M = {}

function M.render(context, opts)
    opts = opts or {}
    local d = opts.destination
    if d and d ~= "" then
        directory.render("contents", context, { destination = d })
    else
        directory.render("contents", context)
    end
    return context
end

return M
